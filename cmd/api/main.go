package main

import (
	"log/slog"
	"net/http"

	"github.com/SAB1825/go_authentication/internal/config"
	"github.com/SAB1825/go_authentication/internal/routes"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:         config.HTTP.Addr,
		ReadTimeout:  config.HTTP.ReadTimeout,
		WriteTimeout: config.HTTP.WriteTimeout,
		IdleTimeout:  config.HTTP.IdleTimeout,
		Handler:      routes.RegisterRoutes(config),
	}

	slog.Info("Starting server", "addr", config.HTTP.Addr)
	server.ListenAndServe()
}
