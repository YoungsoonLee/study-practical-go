package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doSomeWork(data []byte) {
	time.Sleep(15 * time.Second)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	log.Println("ping: Got a request")
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "pong")
}

func handleUserAPI(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)

	log.Println("I started processing the request")

	req, err := http.NewRequestWithContext(r.Context(), "GET", "http://localhost:8080/ping", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	log.Println("Outgoing HTTP request")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error makiing request: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	log.Println("Processing the response i got")

	go func() {
		doSomeWork(data)
		done <- true
	}()

	// waiting two channels
	select {
	case <-done:
		log.Println("doSomeWork done: Continuing request processing")
	case <-r.Context().Done():
		log.Printf("Aborting request processing: %v\n", r.Context().Err())
		return
	}

	fmt.Fprint(w, string(data))
	log.Println("I finished process the request")

	// log.Println(
	// 	"Before continuing, I will check if the timeout has already expire",
	// )
	// if r.Context().Err() != nil {
	// 	log.Printf(
	// 		"Aborting futher processing: %v\n",
	// 		r.Context().Err(),
	// 	)
	// 	return
	// }

	// fmt.Fprintf(w, "Hello world")
	// log.Println("I finished processing the request")
}

func shutDown(ctx context.Context, s *http.Server, waitForShutdownCompletion chan struct{}) {
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigch
	log.Printf("Got signal: %v. Server shutting down.", sig)
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
	waitForShutdownCompletion <- struct{}{}
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	waitForShutdownCompletion := make(chan struct{})
	ctx, cancel := context.WithTimeout(
		context.Background(), 30*time.Second,
	)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users/", handleUserAPI)

	s := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	go shutDown(ctx, &s, waitForShutdownCompletion)

	err := s.ListenAndServe()
	log.Printf(
		"Waiting for shutdown to complete..",
	)
	<-waitForShutdownCompletion
	log.Fatal(err)
}
