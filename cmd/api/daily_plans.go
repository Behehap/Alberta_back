package main

import (
	"errors"
	"net/http"
	"time"

	"strconv"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) createDailyPlanHandler(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	var input struct {
		PlanDate string `json:"plan_date" validate:"required"`
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

	planDate, err := time.Parse("2006-01-02", input.PlanDate)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid plan_date format, please use YYYY-MM-DD"))
		return
	}

	dp := &store.DailyPlan{
		WeeklyPlanID: weeklyPlan.ID,
		PlanDate:     planDate,
	}

	err = app.store.DailyPlans.Insert(r.Context(), dp)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"daily_plan": dp}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listDailyPlansHandler(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	plans, err := app.store.DailyPlans.GetAllForWeeklyPlan(r.Context(), weeklyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"daily_plans": plans}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getDailyPlanHandler(w http.ResponseWriter, r *http.Request) {

	dailyPlan, ok := r.Context().Value(dailyPlanContextKey).(*store.DailyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve daily plan from context"))
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"daily_plan": dailyPlan}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteDailyPlanHandler(w http.ResponseWriter, r *http.Request) {
	dailyPlanID, err := strconv.ParseInt(chi.URLParam(r, "dailyPlanID"), 10, 64)
	if err != nil || dailyPlanID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.store.DailyPlans.Delete(r.Context(), dailyPlanID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "daily plan successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
