package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

// NewLimiter(r, b) returns a new limiter that allows events up to rate r and permits bursts of at most b tokens.
func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4) // 2 requests per second, with a burst of 4 requests

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() { // If the limiter does not allow the request
			message := Message{
				Status: "Request Failed",
				Body:   "The API is at capacity, try again later.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		} else {
			//The anonymous function uses the limiter to check if this request is within the rate limits with limiter.Allow().
			// If it is, the anonymous function calls the next function in the chain i.e endpointHandle()
			// If the request is out of limits, the anonymous function returns an error message to the client.
			next(w, r)
		}
	})

}
