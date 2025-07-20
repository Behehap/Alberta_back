package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

type WeeklyCalendarResponse struct {
	WeeklyPlan     *store.WeeklyPlan    `json:"weekly_plan"`
	DailySchedules []DailyCalendarEntry `json:"daily_schedules"`
}

type DailyCalendarEntry struct {
	DailyPlan     *store.DailyPlan     `json:"daily_plan"`
	StudySessions []StudySessionDetail `json:"study_sessions"`
}

type StudySessionDetail struct {
	ID             int64       `json:"id"`
	DailyPlanID    int64       `json:"daily_plan_id"`
	Book           *store.Book `json:"book"`
	IsCompleted    bool        `json:"is_completed"`
	CompletionDate *time.Time  `json:"completion_date,omitempty"` // Changed to pointer
	StartTime      string      `json:"start_time"`
	EndTime        string      `json:"end_time"`
}

func mapStudySessionToDetail(ss *store.StudySession, book *store.Book) StudySessionDetail {
	var completionDate *time.Time
	if ss.CompletionDate.Valid {
		completionDate = &ss.CompletionDate.Time
	}

	return StudySessionDetail{
		ID:             ss.ID,
		DailyPlanID:    ss.DailyPlanID,
		Book:           book,
		IsCompleted:    ss.IsCompleted,
		CompletionDate: completionDate, // Assign pointer
		StartTime:      ss.StartTime,
		EndTime:        ss.EndTime,
	}
}

func (app *application) createWeeklyPlanHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	var input struct {
		StartDateOfWeek          string `json:"start_date_of_week" validate:"required"`
		DayStartTime             string `json:"day_start_time"`
		MaxStudyTimeHoursPerWeek int    `json:"max_study_time_hours_per_week"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDateOfWeek)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid date format for start_date_of_week, please use YYYY-MM-DD"))
		return
	}

	var dayStartTime sql.NullString
	if input.DayStartTime != "" {
		dayStartTime.String = input.DayStartTime + ":00"
		dayStartTime.Valid = true
	}

	wp := &store.WeeklyPlan{
		StudentID:                student.ID,
		StartDateOfWeek:          startDate,
		DayStartTime:             dayStartTime,
		MaxStudyTimeHoursPerWeek: input.MaxStudyTimeHoursPerWeek,
	}

	err = app.store.WeeklyPlans.Insert(r.Context(), wp)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weekly_plan": wp}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listWeeklyPlansHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	plans, err := app.store.WeeklyPlans.GetAllForStudent(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weekly_plans": plans}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getFullWeeklyCalendarHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	weeklyPlanID, err := strconv.ParseInt(chi.URLParam(r, "planID"), 10, 64) // Corrected from "weeklyPlanID" to "planID"
	if err != nil || weeklyPlanID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	weeklyPlan, err := app.store.WeeklyPlans.Get(r.Context(), weeklyPlanID)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if weeklyPlan.StudentID != student.ID {
		app.notFoundResponse(w, r)
		return
	}

	dailyPlans, err := app.store.DailyPlans.GetAllForWeeklyPlan(r.Context(), weeklyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var dailySchedules []DailyCalendarEntry
	for _, dp := range dailyPlans {
		studySessions, err := app.store.StudySessions.GetAllForDailyPlan(r.Context(), dp.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		var detailedSessions []StudySessionDetail
		for _, ss := range studySessions {
			book, err := app.store.Books.Get(r.Context(), ss.BookID)
			if err != nil {
				app.logger.Printf("Warning: Could not retrieve book %d for study session %d: %v", ss.BookID, ss.ID, err)
				detailedSessions = append(detailedSessions, mapStudySessionToDetail(ss, nil))
				continue
			}
			detailedSessions = append(detailedSessions, mapStudySessionToDetail(ss, book))
		}
		dailySchedules = append(dailySchedules, DailyCalendarEntry{
			DailyPlan:     dp,
			StudySessions: detailedSessions,
		})
	}

	response := WeeklyCalendarResponse{
		WeeklyPlan:     weeklyPlan,
		DailySchedules: dailySchedules,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weekly_calendar": response}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
