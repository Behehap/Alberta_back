// cmd/api/api.go
package main

import (
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

// mount sets up the router and all the API routes.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// All API routes will be under the /v1 prefix
	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

		// Routes for listing seeded data like grades, majors, and books.
		r.Get("/grades", app.listGradesHandler)
		r.Get("/majors", app.listMajorsHandler)
		r.Get("/curriculum/books", app.listBooksForCurriculumHandler)

		// Routes for managing students.
		r.Post("/students", app.createStudentHandler)

		// This group handles all routes related to a specific student.
		r.Route("/students/{studentID}", func(r chi.Router) {
			// This middleware fetches the student for all routes in this group.
			r.Use(app.studentContextMiddleware)

			// Endpoints for the student resource itself.
			r.Get("/", app.getStudentHandler)
			r.Patch("/", app.updateStudentHandler)
			r.Delete("/", app.deleteStudentHandler)

			// Endpoints for a student's unavailable times.
			r.Get("/unavailable-times", app.listUnavailableTimesHandler)
			r.Post("/unavailable-times", app.createUnavailableTimeHandler)

			// Endpoints for a student's weekly plans.
			r.Get("/weekly-plans", app.listWeeklyPlansHandler)
			r.Post("/weekly-plans", app.createWeeklyPlanHandler)
		})
	})

	return r
}
