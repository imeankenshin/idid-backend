package settings

import (
	"idid-api/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Middlewares is a slice of middleware functions
var Middlewares = []func(http.Handler) http.Handler{
	middleware.Logger,
	middleware.Recoverer,
}

// ApplicationRoutes creates a REST router for the posts resource
func ApplicationRoutes(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/posts", controllers.PostControllers)
}
