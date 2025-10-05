package main

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Behehap/Alberta/internal/store"
)

type FrequencyCalculationRequest struct {
	SelectedSubjects []int64 `json:"selected_subjects" validate:"required,min=1,max=20"`
	TemplateID       *int64  `json:"template_id,omitempty"`
}

type FrequencyCalculationResponse struct {
	RecommendedTemplate *store.ScheduleTemplate `json:"recommended_template"`
	AdjustedFrequencies map[int64]int           `json:"adjusted_frequencies"`
	TotalWeeklyBlocks   int                     `json:"total_weekly_blocks"`
}

func (app *application) calculateFrequenciesHandler(w http.ResponseWriter, r *http.Request) {
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

	var input FrequencyCalculationRequest
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

	// Calculate total weekly blocks (daily hours * 6 days, excluding Friday)
	totalWeeklyBlocks := calculateTotalWeeklyBlocks(weeklyPlan)

	var selectedTemplate *store.ScheduleTemplate

	if input.TemplateID != nil {
		// Use specified template
		selectedTemplate, err = app.store.ScheduleTemplates.Get(r.Context(), *input.TemplateID)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("specified template not found"))
			return
		}
	} else {
		// Find closest template automatically
		selectedTemplate, err = app.scheduler.FindClosestTemplate(r.Context(), student.GradeID, student.MajorID, totalWeeklyBlocks)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	if selectedTemplate == nil {
		app.badRequestResponse(w, r, errors.New("no suitable template found for your grade and major"))
		return
	}

	// Calculate adjusted frequencies
	adjustedFrequencies, err := app.scheduler.CalculateAdjustedFrequencies(
		r.Context(),
		selectedTemplate,
		input.SelectedSubjects,
		totalWeeklyBlocks,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := FrequencyCalculationResponse{
		RecommendedTemplate: selectedTemplate,
		AdjustedFrequencies: adjustedFrequencies,
		TotalWeeklyBlocks:   totalWeeklyBlocks,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"frequency_calculation": response}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// calculateTotalWeeklyBlocks calculates total study blocks for the week
func calculateTotalWeeklyBlocks(weeklyPlan *store.WeeklyPlan) int {
	// Assuming 100-minute blocks (as in current scheduler)
	// Convert max study time hours to blocks
	if weeklyPlan.MaxStudyTimeHoursPerWeek > 0 {
		// Convert hours to 100-minute blocks
		totalMinutes := weeklyPlan.MaxStudyTimeHoursPerWeek * 60
		return totalMinutes / 100
	}

	// Default calculation if no max study time specified
	// This should use day_start_time and day_end_time to calculate available slots
	// For now, return a reasonable default
	return 18 // Default to 18 blocks (3 hours/day * 6 days)
}

func (app *application) getRecommendedTemplateHandler(w http.ResponseWriter, r *http.Request) {
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

	// Get selected subjects from query parameter
	selectedSubjectsParam := r.URL.Query().Get("selected_subjects")
	var selectedBookIDs []int64

	if selectedSubjectsParam != "" {
		// Parse comma-separated book IDs
		subjectStrs := strings.Split(selectedSubjectsParam, ",")
		for _, str := range subjectStrs {
			if id, err := strconv.ParseInt(str, 10, 64); err == nil {
				selectedBookIDs = append(selectedBookIDs, id)
			}
		}
	}

	totalWeeklyBlocks := calculateTotalWeeklyBlocks(weeklyPlan)

	recommendedTemplate, err := app.scheduler.FindClosestTemplate(r.Context(), student.GradeID, student.MajorID, totalWeeklyBlocks)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if recommendedTemplate == nil {
		app.writeJSON(w, http.StatusOK, envelope{"recommended_template": nil}, nil)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"recommended_template": recommendedTemplate}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
