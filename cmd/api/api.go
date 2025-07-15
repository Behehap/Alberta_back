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

type contextKey string

const studentContextKey = contextKey("student")
const weeklyPlanContextKey = contextKey("weekly_plan")
const dailyPlanContextKey = contextKey("daily_plan")
const studySessionContextKey = contextKey("study_session")
const examScheduleContextKey = contextKey("exam_schedule")

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

type application struct {
	config config
	logger *log.Logger
	store  *store.Storage
}

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

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)

		r.Get("/grades", app.listGradesHandler)
		r.Get("/majors", app.listMajorsHandler)
		r.Get("/curriculum/books", app.listBooksForCurriculumHandler)

		r.Post("/exam-schedules", app.createExamScheduleHandler)
		r.Route("/exam-schedules/{examID}", func(r chi.Router) {
			r.Use(app.examScheduleContextMiddleware)
			r.Get("/", app.getExamScheduleHandler)
			r.Get("/scope", app.listExamScopeItemsHandler)
			r.Post("/scope", app.createExamScopeItemHandler)
		})

		r.Post("/students", app.createStudentHandler)
		r.Route("/students/{studentID}", func(r chi.Router) {
			r.Use(app.studentContextMiddleware)
			r.Get("/", app.getStudentHandler)
			r.Patch("/", app.updateStudentHandler)
			r.Delete("/", app.deleteStudentHandler)

			r.Get("/exam-schedules", app.listExamSchedulesHandler)

			r.Get("/unavailable-times", app.listUnavailableTimesHandler)
			r.Post("/unavailable-times", app.createUnavailableTimeHandler)

			r.Get("/weekly-plans", app.listWeeklyPlansHandler)
			r.Post("/weekly-plans", app.createWeeklyPlanHandler)

			r.Route("/weekly-plans/{planID}", func(r chi.Router) {
				r.Use(app.weeklyPlanContextMiddleware)

				r.Get("/subject-frequencies", app.listSubjectFrequenciesHandler)
				r.Post("/subject-frequencies", app.createSubjectFrequencyHandler)

				r.Get("/daily-plans", app.listDailyPlansHandler)
				r.Post("/daily-plans", app.createDailyPlanHandler)

				r.Route("/daily-plans/{dailyPlanID}", func(r chi.Router) {
					r.Use(app.dailyPlanContextMiddleware)

					r.Get("/study-sessions", app.listStudySessionsHandler)
					r.Post("/study-sessions", app.createStudySessionHandler)

					r.Route("/study-sessions/{sessionID}", func(r chi.Router) {
						r.Use(app.studySessionContextMiddleware)

						r.Get("/report", app.getSessionReportHandler)
						r.Post("/report", app.createSessionReportHandler)
					})
				})
			})
		})
	})

	return r
}
