package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/h-roth/sequence-service/sequence"
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

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)
	http.HandleFunc("/cmd/client", func(w http.ResponseWriter, r *http.Request) {
		request()
	})

	server := newServer(withPort(*port))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	nextID := sequence.PublicGetNextID()
	fmt.Fprintf(w, "%d", nextID)
}
