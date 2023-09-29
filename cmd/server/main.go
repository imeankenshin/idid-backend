package main

import (
	config "idid-api/internal/settings"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(config.Middlewares...)
	r.Group(config.ApplicationRoutes)
	log.Fatalln(http.ListenAndServe(":4000", r))
}
