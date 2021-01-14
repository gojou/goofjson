package handlers

import (
	"fmt"
	"net/http"
)

// Home is where I want to be
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Would you like to play a game?")
}
