package main

import (
	"fmt"
	"io"
	"net/http"
)

// NOTE: Prefer using a load testing client like `ali` for this.
// Refer to README.md for more information.

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

	for i := 0; i < numRequests; i++ {
		go func(num int) {
			results <- mockHTTPRequest(num)
		}(i)
	}

	// Receive results from the channel and log.
	for i := 0; i < numRequests; i++ {
		fmt.Println(<-results)
	}
}
