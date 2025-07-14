package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createDailyPlanHandler(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	var input struct {
		PlanDate string `json:"plan_date" validate:"required"` // Expects "YYYY-MM-DD"
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	planDate, err := time.Parse("2006-01-02", input.PlanDate)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid date format for plan_date, please use YYYY-MM-DD"))
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
