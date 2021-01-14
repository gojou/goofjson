package handlers

import (
	"fmt"
	"net/http"
)

// Counter demonstrates a closure by incrementing a counter
// every time the handler is called.
func Counter() func(w http.ResponseWriter, r *http.Request) {
	var c int32
	c = 0

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c = c + 1
		fmt.Fprintf(w, "This is a test: %v\n", c)
	})
}

// This works too. Counter returning an http.HandlerFunc (NOT an http.Handler)
// is returning a data structure that includes http.ResponseWriter and
// *http.Request, so the compiler is fine with it. For me at this point that
// just muddies the fact that for the closure to work the function's input must match
// what is returned.

// // Counter is counting
// func Counter() http.HandlerFunc {
// 	var c int32
// 	c = 0

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		c = c + 1
// 		fmt.Fprintf(w, "This is a test: %v\n", c)
// 	})
// }
