package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createUnavailableTimeHandler(w http.ResponseWriter, r *http.Request) {

	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	var input struct {
		Title       string `json:"title"`
		DayOfWeek   int    `json:"day_of_week" validate:"gte=0,lte=6"`
		StartTime   string `json:"start_time" validate:"required"`
		EndTime     string `json:"end_time" validate:"required"`
		IsRecurring bool   `json:"is_recurring"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(input)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	ut := &store.UnavailableTime{
		StudentID:   student.ID,
		Title:       input.Title,
		DayOfWeek:   input.DayOfWeek,
		StartTime:   input.StartTime + ":00",
		EndTime:     input.EndTime + ":00",
		IsRecurring: input.IsRecurring,
	}

	err = app.store.UnavailableTimes.Insert(r.Context(), ut)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"unavailable_time": ut}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listUnavailableTimesHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	times, err := app.store.UnavailableTimes.GetAllForStudent(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"unavailable_times": times}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
