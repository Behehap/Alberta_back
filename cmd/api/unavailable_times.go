package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

type UnavailableTimeDisplay struct {
	ID          int64  `json:"id"`
	StudentID   int64  `json:"student_id"`
	Title       string `json:"title"`
	DayOfWeek   int    `json:"day_of_week"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	IsRecurring bool   `json:"is_recurring"`
}

func mapUnavailableTimeToDisplay(ut *store.UnavailableTime) *UnavailableTimeDisplay {
	displayUt := &UnavailableTimeDisplay{
		ID:          ut.ID,
		StudentID:   ut.StudentID,
		Title:       ut.Title,
		DayOfWeek:   ut.DayOfWeek,
		StartTime:   ut.StartTime.Format("15:04:05"),
		EndTime:     ut.EndTime.Format("15:04:05"),
		IsRecurring: ut.IsRecurring,
	}

	return displayUt
}

func (app *application) createUnavailableTimeHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	var input struct {
		Title       string `json:"title" validate:"required"`
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

	startTime, err := time.Parse("15:04", input.StartTime)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid start_time format, please use HH:MM"))
		return
	}

	endTime, err := time.Parse("15:04", input.EndTime)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid end_time format, please use HH:MM"))
		return
	}

	if endTime.Before(startTime) {
		app.badRequestResponse(w, r, errors.New("end time cannot be before start time"))
		return
	}

	ut := &store.UnavailableTime{
		StudentID:   student.ID,
		Title:       input.Title,
		DayOfWeek:   input.DayOfWeek,
		StartTime:   startTime,
		EndTime:     endTime,
		IsRecurring: input.IsRecurring,
	}

	err = app.store.UnavailableTimes.Insert(r.Context(), ut)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"unavailable_time": mapUnavailableTimeToDisplay(ut)}, nil)
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

	displayTimes := make([]*UnavailableTimeDisplay, len(times))
	for i, t := range times {
		displayTimes[i] = mapUnavailableTimeToDisplay(t)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"unavailable_times": displayTimes}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateUnavailableTimeHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	utID, err := strconv.ParseInt(chi.URLParam(r, "utID"), 10, 64)
	if err != nil || utID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	ut, err := app.store.UnavailableTimes.Get(r.Context(), utID)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if ut.StudentID != student.ID {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Title       *string `json:"title"`
		DayOfWeek   *int    `json:"day_of_week"`
		StartTime   *string `json:"start_time"`
		EndTime     *string `json:"end_time"`
		IsRecurring *bool   `json:"is_recurring"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		ut.Title = *input.Title
	}

	if input.DayOfWeek != nil {
		if *input.DayOfWeek < 0 || *input.DayOfWeek > 6 {
			app.badRequestResponse(w, r, errors.New("day_of_week must be between 0 and 6"))
			return
		}
		ut.DayOfWeek = *input.DayOfWeek
	}

	if input.StartTime != nil {
		parsedTime, err := time.Parse("15:04", *input.StartTime)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("invalid start_time format, please use HH:MM"))
			return
		}
		ut.StartTime = parsedTime
	}

	if input.EndTime != nil {
		parsedTime, err := time.Parse("15:04", *input.EndTime)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("invalid end_time format, please use HH:MM"))
			return
		}
		ut.EndTime = parsedTime
	}

	if input.IsRecurring != nil {
		ut.IsRecurring = *input.IsRecurring
	}

	err = app.store.UnavailableTimes.Update(r.Context(), ut)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"unavailable_time": mapUnavailableTimeToDisplay(ut)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteUnavailableTimeHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	utID, err := strconv.ParseInt(chi.URLParam(r, "utID"), 10, 64)
	if err != nil || utID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	ut, err := app.store.UnavailableTimes.Get(r.Context(), utID)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	if ut.StudentID != student.ID {
		app.notFoundResponse(w, r)
		return
	}

	err = app.store.UnavailableTimes.Delete(r.Context(), utID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "unavailable time deleted successfully"}, nil)
}
