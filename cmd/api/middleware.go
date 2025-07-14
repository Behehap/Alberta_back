// cmd/api/middleware.go
package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) studentContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		studentID, err := strconv.ParseInt(chi.URLParam(r, "studentID"), 10, 64)
		if err != nil || studentID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		student, err := app.store.Students.Get(r.Context(), studentID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), studentContextKey, student)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// This new middleware fetches a weekly plan by its ID.
func (app *application) weeklyPlanContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		planID, err := strconv.ParseInt(chi.URLParam(r, "planID"), 10, 64)
		if err != nil || planID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		plan, err := app.store.WeeklyPlans.Get(r.Context(), planID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), weeklyPlanContextKey, plan)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
