package main

import (
	"fmt"
	"idid-api/internal/settings"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(settings.Middlewares...)
	r.Group(settings.ApplicationRoutes)
	fmt.Println("Server running on http://localhost:4000")
	log.Fatalln(http.ListenAndServe(":4000", r))
}
