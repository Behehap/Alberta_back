package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) generateWeeklyScheduleHandler(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		ScheduleTemplateID int64 `json:"schedule_template_id" validate:"required,gt=0"`
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

	selectedTemplate, err := app.store.ScheduleTemplates.Get(r.Context(), input.ScheduleTemplateID)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			app.badRequestResponse(w, r, errors.New("selected schedule template not found"))
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	totalStudyBlocks := selectedTemplate.TotalStudyBlocksPerWeek

	templateRules, err := app.store.TemplateRules.GetAllForTemplate(r.Context(), selectedTemplate.ID)
	if err != nil {
		app.logger.Printf("Could not retrieve template rules for template %d: %v", selectedTemplate.ID, err)
		templateRules = []*store.TemplateRule{}
	}

	subjectFrequencies, err := app.store.SubjectFrequencies.GetAllForWeeklyPlan(r.Context(), weeklyPlan.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	unavailableTimes, err := app.store.UnavailableTimes.GetAllForStudent(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.scheduler.GenerateWeeklyPlan(
		r.Context(),
		student.ID,
		weeklyPlan.ID,
		weeklyPlan.StartDateOfWeek,
		totalStudyBlocks,
		unavailableTimes,
		subjectFrequencies,
		templateRules,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "weekly schedule generated successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
