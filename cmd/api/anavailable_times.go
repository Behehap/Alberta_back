// cmd/api/unavailable_times.go
package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

// Handles POST /v1/students/{studentID}/unavailable-times
func (app *application) createUnavailableTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Grab the student from the context, which the middleware already got for us.
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	// A simple struct to read the incoming JSON into.
	var input struct {
		Title       string `json:"title"`
		DayOfWeek   int    `json:"day_of_week" validate:"gte=0,lte=6"`
		StartTime   string `json:"start_time" validate:"required"` // Expecting "HH:MM"
		EndTime     string `json:"end_time" validate:"required"`   // Expecting "HH:MM"
		IsRecurring bool   `json:"is_recurring"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Quick validation on the input.
	err = Validate.Struct(input)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Map the validated input to our actual data model.
	ut := &store.UnavailableTime{
		StudentID:   student.ID,
		Title:       input.Title,
		DayOfWeek:   input.DayOfWeek,
		StartTime:   input.StartTime + ":00", // Append seconds for the DB TIME type.
		EndTime:     input.EndTime + ":00",
		IsRecurring: input.IsRecurring,
	}

	// Pass it to the data layer to save.
	err = app.store.UnavailableTimes.Insert(r.Context(), ut)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send back a "201 Created" response.
	err = app.writeJSON(w, http.StatusCreated, envelope{"unavailable_time": ut}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handles GET /v1/students/{studentID}/unavailable-times
func (app *application) listUnavailableTimesHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	// Ask the data layer for the list of times.
	times, err := app.store.UnavailableTimes.GetAllForStudent(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send the list back to the client.
	err = app.writeJSON(w, http.StatusOK, envelope{"unavailable_times": times}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
