package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type requestContextKey struct{}
type requestContextValue struct {
	requestID string
}

type logLine struct {
	UserIP string `json:"user_ip"`
	Event  string `json:"event"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)

	for {
		var l logLine
		err := dec.Decode(&l)
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(l.UserIP, l.Event)
	}
	fmt.Fprintf(w, "OK")
}

func addRequestID(r *http.Request, requestID string) *http.Request {
	c := requestContextValue{
		requestID: requestID,
	}

	currentCtx := r.Context()
	newCtx := context.WithValue(currentCtx, requestContextKey{}, c)
	return r.WithContext(newCtx)
}

func apiHandler(w http.ResponseWriter, r *http.Request, config appConfig) {
	//requestID := "request-123-abc"
	//r = addRequestID(r, requestID)
	//processRequest(w, r)
	config.logger.Println("Handling API request")
	fmt.Fprintf(w, "Hello, world!")
}

func logRequest(r *http.Request) {
	ctx := r.Context()
	v := ctx.Value(requestContextKey{})

	if m, ok := v.(requestContextValue); ok {
		log.Printf("Processing request: %s", m.requestID)
	}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprintf(w, "Request processed")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request, config appConfig) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	config.logger.Println("Handling healthcheck request")
	fmt.Fprintf(w, "ok")
}

func longRunningProcess(logWriter *io.PipeWriter) {
	for i := 0; i <= 20; i++ {
		fmt.Fprintf(
			logWriter,
			`{"id":%d, "user_ip": "172.121.19.21", "event": "click_on_add_cart"}`, i,
		)
		fmt.Fprintln(logWriter)
		time.Sleep(1 * time.Second)
	}
	logWriter.Close()
}

func longRunningProcessHandler(w http.ResponseWriter, r *http.Request) {
	done := make(chan struct{})
	logReader, logWriter := io.Pipe()
	go longRunningProcess(logWriter)
	go progressStreamer(logReader, w, done)

	<-done
}

func progressStreamer(logReader *io.PipeReader, w http.ResponseWriter, done chan struct{}) {
	buf := make([]byte, 500)

	f, flushSupported := w.(http.Flusher)

	defer logReader.Close()

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	for {
		n, err := logReader.Read(buf)
		if err == io.EOF {
			break
		}
		w.Write(buf[:n])
		if flushSupported {
			f.Flush()
		}
	}
	done <- struct{}{}
}

func newConfig() appConfig {
	config := appConfig{
		logger: log.New(
			os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile,
		),
	}
	return config
}

func setupHandlers(mux *http.ServeMux, config appConfig) {

	mux.Handle("/healthz", &app{config: config, handler: healthCheckHandler})
	mux.Handle("/api", &app{config: config, handler: apiHandler})
	mux.HandleFunc("/decode", decodeHandler)
	mux.HandleFunc("/job", longRunningProcessHandler)
}

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	c := newConfig()

	mux := http.NewServeMux()
	setupHandlers(mux, c)

	m := loggingMiddleware(panicMiddleware(mux)) // !!!changing middleware

	log.Fatal(http.ListenAndServe(listenAddr, m))
}
