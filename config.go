package main

import (
	"log"
	"net/http"
	"time"
)

type config struct {
	port string
	id   int64
}
type configFunc func(*config)

func defaultConfig() *config {
	return &config{
		port: "8080",
		id:   0,
	}
}

func withPort(port string) configFunc {
	return func(c *config) {
		c.port = port
	}
}

// Slice of opt func, allows us to give 1 or many configs options
// Advantagous as it provides clear defaults but is also easily overridden
func newServer(opts ...configFunc) *http.Server {
	c := defaultConfig()
	for _, opt := range opts {
		opt(c)
	}

	log.Printf("Starting server with config: %+v\n", c)

	return &http.Server{
		Addr:              ":" + c.port,
		ReadHeaderTimeout: 3 * time.Second,
	}
}
