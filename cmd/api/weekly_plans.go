package main

import (
	"database/sql" // We need to import this to use sql.NullString
	"errors"
	"net/http"
	"time"

	"github.com/Behehap/Alberta/internal/store"
)

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

	// This is the fix. We check if the start time was provided.
	// If not, we use a special type that will be saved as NULL in the database.
	var dayStartTime sql.NullString
	if input.DayStartTime != "" {
		dayStartTime.String = input.DayStartTime + ":00"
		dayStartTime.Valid = true
	}

	wp := &store.WeeklyPlan{
		StudentID:                student.ID,
		StartDateOfWeek:          startDate,
		DayStartTime:             dayStartTime, // Use the new nullable string
		MaxStudyTimeHoursPerWeek: input.MaxStudyTimeHoursPerWeek,
	}

	// We also need to update the data layer to handle this new type.
	err = app.store.WeeklyPlans.Insert(r.Context(), wp)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"weekly_plan": wp}, nil)
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
