// cmd/api/middleware.go
package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Behehap/Alberta/internal/store"
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

func (app *application) weeklyPlanContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// First, get the student from the context.
		// This middleware MUST run after the studentContextMiddleware.
		student, ok := r.Context().Value(studentContextKey).(*store.Student)
		if !ok {
			app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
			return
		}

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

		// *** THIS IS THE CRITICAL FIX ***
		// Check if the plan's student ID matches the student ID from the context.
		if plan.StudentID != student.ID {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), weeklyPlanContextKey, plan)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) weeklyStudyItemContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemID, err := strconv.ParseInt(chi.URLParam(r, "itemID"), 10, 64)
		if err != nil || itemID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		item, err := app.store.WeeklyStudyItems.Get(r.Context(), itemID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), weeklyStudyItemContextKey, item)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) examScheduleContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		examID, err := strconv.ParseInt(chi.URLParam(r, "examID"), 10, 64)
		if err != nil || examID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		exam, err := app.store.ExamSchedules.Get(r.Context(), examID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), examScheduleContextKey, exam)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
