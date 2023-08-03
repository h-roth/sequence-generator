package main

import (
	"fmt"
	"io"
	"net/http"
)

func mockHTTPRequest(num int) string {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/?request=%d", num))
	if err != nil {
		return fmt.Sprintf("Request %d failed with error: %s", num, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Request %d failed when reading body with error: %s", num, err)
	}

	return fmt.Sprintf("Request %d: %s", num, string(body))
}

func request() {
	numRequests := 10000
	results := make(chan string)

	// Use goroutine to send multiple time-consuming jobs to the channel.
	for i := 0; i < numRequests; i++ {
		go func(num int) {
			results <- mockHTTPRequest(num)
		}(i)
	}

	// Receive results from the channel and use them.
	for i := 0; i < numRequests; i++ {
		fmt.Println(<-results)
	}
}
