package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

type contextKey string

const studentContextKey = contextKey("student")
const weeklyPlanContextKey = contextKey("weekly_plan")
const dailyPlanContextKey = contextKey("daily_plan")
const studySessionContextKey = contextKey("study_session")
const examScheduleContextKey = contextKey("exam_schedule")
const scheduleTemplateContextKey = contextKey("schedule_template")
const weeklyStudyItemContextKey = contextKey("weekly_study_item")

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

		if plan.StudentID != student.ID {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), weeklyPlanContextKey, plan)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) dailyPlanContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dailyPlanID, err := strconv.ParseInt(chi.URLParam(r, "dailyPlanID"), 10, 64)
		if err != nil || dailyPlanID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		dailyPlan, err := app.store.DailyPlans.Get(r.Context(), dailyPlanID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), dailyPlanContextKey, dailyPlan)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) studySessionContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := strconv.ParseInt(chi.URLParam(r, "sessionID"), 10, 64)
		if err != nil || sessionID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		session, err := app.store.StudySessions.Get(r.Context(), sessionID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), studySessionContextKey, session)
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

func (app *application) scheduleTemplateContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templateID, err := strconv.ParseInt(chi.URLParam(r, "templateID"), 10, 64)
		if err != nil || templateID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		template, err := app.store.ScheduleTemplates.Get(r.Context(), templateID)
		if err != nil {
			app.notFoundResponse(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), scheduleTemplateContextKey, template)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
