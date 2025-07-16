package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) addWeeklyStudyItemHandler(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	var input struct {
		LessonID int64 `json:"lesson_id" validate:"required,gt=0"`
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

	wsi := &store.WeeklyStudyItem{
		WeeklyPlanID: weeklyPlan.ID,
		LessonID:     input.LessonID,
	}

	err = app.store.WeeklyStudyItems.Insert(r.Context(), wsi)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"weekly_study_item": wsi}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listWeeklyStudyItemsHandler(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	items, err := app.store.WeeklyStudyItems.GetAllForWeeklyPlan(r.Context(), weeklyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weekly_study_items": items}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateWeeklyStudyItemHandler(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		IsCompleted *bool `json:"is_completed"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.IsCompleted != nil {
		item.IsCompleted = *input.IsCompleted
	}

	err = app.store.WeeklyStudyItems.Update(r.Context(), item)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weekly_study_item": item}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
