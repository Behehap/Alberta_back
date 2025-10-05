package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

type ScheduleGenerationRequest struct {
	ScheduleTemplateID *int64        `json:"schedule_template_id,omitempty"`
	SubjectFrequencies map[int64]int `json:"subject_frequencies" validate:"required,min=1"`
}

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

	var input ScheduleGenerationRequest
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

	// Calculate total study blocks from frequencies
	totalStudyBlocks := 0
	for _, freq := range input.SubjectFrequencies {
		totalStudyBlocks += freq
	}

	// Convert frequency map to SubjectFrequency slice for the scheduler
	var subjectFrequencies []*store.SubjectFrequency
	for bookID, frequency := range input.SubjectFrequencies {
		sf := &store.SubjectFrequency{
			WeeklyPlanID:     weeklyPlan.ID,
			BookID:           bookID,
			FrequencyPerWeek: frequency,
		}
		subjectFrequencies = append(subjectFrequencies, sf)
	}

	// Get template rules if template ID is provided
	var templateRules []*store.TemplateRule
	if input.ScheduleTemplateID != nil {
		templateRules, err = app.store.TemplateRules.GetAllForTemplate(r.Context(), *input.ScheduleTemplateID)
		if err != nil {
			app.logger.Printf("Could not retrieve template rules for template %d: %v", *input.ScheduleTemplateID, err)
			templateRules = []*store.TemplateRule{}
		}
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
