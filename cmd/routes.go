package main

import (
	"github.com/gojou/goofjson/pkg/handlers"
	"github.com/gorilla/mux"
)

// Routing blah blah blah
func routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/home", handlers.Home)
	r.HandleFunc("/counter", handlers.Counter())
}
