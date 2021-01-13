package routes

import (
	"github.com/gojou/goofjson/pkg/handlers"
	"github.com/gorilla/mux"
)

// Routing blah blah blah
func Routing(r *mux.Router) {
	r.HandleFunc("/", handlers.Home)
}
