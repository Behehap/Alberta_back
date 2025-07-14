// cmd/api/api.go
package main

import (
	// We need to import context now
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// A custom key type to use for our context. This is a standard Go trick
// to avoid key collisions in the context.
type contextKey string

const studentContextKey = contextKey("student")

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

		// Student routes
		r.Post("/students", app.createStudentHandler)

		// This is a route group for a single student.
		// Any routes inside here will have the /v1/students/{studentID} prefix.
		r.Route("/students/{studentID}", func(r chi.Router) {
			// Use our middleware to fetch the student and put it in the context.
			r.Use(app.studentContextMiddleware)
			r.Get("/", app.getStudentHandler)
			// Adding the update and delete routes.
			r.Patch("/", app.updateStudentHandler)
			r.Delete("/", app.deleteStudentHandler)
		})
	})

	return r
}
