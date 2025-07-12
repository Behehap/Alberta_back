// cmd/api/api.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// Make sure this module path matches your go.mod file

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// config struct holds all the configuration settings for our application.
type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

// application struct holds the dependencies for our HTTP handlers, helpers, and middleware.
type application struct {
	config config
	logger *log.Logger
	store  *store.Storage
}

// run starts the HTTP server.
func (app *application) run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.mount(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Printf("starting %s server on %s", app.config.env, srv.Addr)
	return srv.ListenAndServe()
}

// mount sets up the router and middleware.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger) // Chi's logger
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// API routes
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

		// We will add student routes, schedule routes, etc. here later.
	})

	return r
}
