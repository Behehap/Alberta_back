package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	// Get the daily plan from the context, which our middleware provides.
	dailyPlan, ok := r.Context().Value(dailyPlanContextKey).(*store.DailyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve daily plan from context"))
		return
	}

	var input struct {
		LessonID  int64  `json:"lesson_id" validate:"required,gt=0"`
		StartTime string `json:"start_time" validate:"required"`
		EndTime   string `json:"end_time" validate:"required"`
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

	ss := &store.StudySession{
		DailyPlanID: dailyPlan.ID,
		LessonID:    input.LessonID,
		StartTime:   input.StartTime + ":00",
		EndTime:     input.EndTime + ":00",
	}

	err = app.store.StudySessions.Insert(r.Context(), ss)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"study_session": ss}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listStudySessionsHandler(w http.ResponseWriter, r *http.Request) {

	dailyPlan, ok := r.Context().Value(dailyPlanContextKey).(*store.DailyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve daily plan from context"))
		return
	}

	sessions, err := app.store.StudySessions.GetAllForDailyPlan(r.Context(), dailyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"study_sessions": sessions}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
