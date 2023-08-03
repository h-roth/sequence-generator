package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/h-roth/sequence-service/sequence"
)

// Could move this within config, think it's fine here for now
var generator sequence.Sequence

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	// Interesting note:
	// Originally had this as just "/", this led to browser doing preflight OPTIONS request and accidentally incrementing the sequence
	http.HandleFunc("/sequence", handler)
	http.HandleFunc("/cmd/client", func(w http.ResponseWriter, r *http.Request) {
		request()
	})

	generator = *sequence.New(&sequence.TimestampGenerator{})

	server := newServer(withPort(*port))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Only GET requests are supported")
		return
	}

	nextID := generator.GetNextID()
	log.Printf("Next ID: %d", nextID)
	fmt.Fprintf(w, "%d", nextID)
}
