package routes

import (
	"github.com/gojou/bones/pkg/handlers"
	"github.com/gorilla/mux"
)

// Routing blah blah blah
func Routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
}
