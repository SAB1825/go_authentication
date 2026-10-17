package routes

import (
	"github.com/SAB1825/go_authentication/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type IRoutes interface {
	RegisterRoutes(router *chi.Mux)
}

func RegisterRoutes(config *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	HealthRoute(r)

	return r
}
