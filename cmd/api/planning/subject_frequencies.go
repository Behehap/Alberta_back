package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createSubjectFrequencyHandler(w http.ResponseWriter, r *http.Request) {

	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	var input struct {
		BookID           int64 `json:"book_id" validate:"required,gt=0"`
		FrequencyPerWeek int   `json:"frequency_per_week" validate:"required,gt=0"`
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

	sf := &store.SubjectFrequency{
		WeeklyPlanID:     weeklyPlan.ID,
		BookID:           input.BookID,
		FrequencyPerWeek: input.FrequencyPerWeek,
	}

	err = app.store.SubjectFrequencies.Insert(r.Context(), sf)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"subject_frequency": sf}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listSubjectFrequenciesHandler(w http.ResponseWriter, r *http.Request) {

	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	frequencies, err := app.store.SubjectFrequencies.GetAllForWeeklyPlan(r.Context(), weeklyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"subject_frequencies": frequencies}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
