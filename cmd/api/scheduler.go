package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/scheduler"
	"github.com/Behehap/Alberta/internal/store"
)

// This handler will trigger our scheduling algorithm.
func (app *application) generateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// Get the student and weekly plan from the context.
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	weeklyPlan, ok := r.Context().Value(weeklyPlanContextKey).(*store.WeeklyPlan)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly plan from context"))
		return
	}

	// The user needs to tell us which template to use for the generation.
	var input struct {
		TemplateID int64 `json:"template_id" validate:"required,gt=0"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Create a new instance of our scheduler.
	sched := scheduler.New(app.store)

	// Call the Generate method. This is where all the magic happens.
	err = sched.Generate(r.Context(), student.ID, weeklyPlan.ID, input.TemplateID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// If everything went well, send back a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "schedule generated successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
