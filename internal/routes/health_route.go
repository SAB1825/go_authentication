package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthRoute(r chi.Router) {
	r.Route("/api/v1/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"ok"}`))
		})
	})
}
