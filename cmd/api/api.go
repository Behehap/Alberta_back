package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// config and application structs are defined ONLY in cmd/api/main.go
// contextKey type and all const context key declarations are defined ONLY in cmd/api/middleware.go

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
		r.Get("/books/{bookID}/lessons", app.listLessonsForBookHandler)

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

				r.Post("/generate", app.generateWeeklyScheduleHandler)
				r.Get("/calendar", app.getFullWeeklyCalendarHandler)

				r.Get("/subject-frequencies", app.listSubjectFrequenciesHandler)
				r.Post("/subject-frequencies", app.createSubjectFrequencyHandler)

				r.Get("/daily-plans", app.listDailyPlansHandler)
				r.Post("/daily-plans", app.createDailyPlanHandler)
			})
		})

		// Corrected to use Study Sessions terminology and correct middleware/handlers
		r.Route("/daily-plans/{dailyPlanID}", func(r chi.Router) {
			r.Use(app.dailyPlanContextMiddleware)
			r.Get("/", app.getDailyPlanHandler)
			r.Get("/study-sessions", app.listStudySessionsHandler)
			r.Post("/study-sessions", app.createStudySessionHandler)
		})

		r.Route("/study-sessions/{sessionID}", func(r chi.Router) {
			r.Use(app.studySessionContextMiddleware)
			r.Get("/", app.getStudySessionHandler)
			r.Patch("/", app.updateStudySessionHandler)
			r.Delete("/", app.deleteStudySessionHandler)

			r.Get("/report", app.getSessionReportHandler)
			r.Post("/report", app.createSessionReportHandler)
		})
	})

	return r
}
