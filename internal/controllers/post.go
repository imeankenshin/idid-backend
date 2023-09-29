package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// PostControllers is a REST controller for the posts resource
func PostControllers(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts index"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts create"))
	})
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts show"))
	})
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts update"))
	})
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts delete"))
	})
}
