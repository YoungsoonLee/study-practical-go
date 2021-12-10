package main

import (
	"log"
	"net/http"
	"time"
)

type appConfig struct {
	logger *log.Logger
}

type app struct {
	config  appConfig
	handler func(w http.ResponseWriter, r *http.Request, config appConfig)
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	a.handler(w, r, a.config)
	a.config.logger.Printf(
		"path=%s method=%s duration=%f", r.URL, r.Method, time.Now().Sub(startTime).Seconds(),
	)
}
