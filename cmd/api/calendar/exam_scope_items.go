package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createExamScopeItemHandler(w http.ResponseWriter, r *http.Request) {

	exam, ok := r.Context().Value(examScheduleContextKey).(*store.ExamSchedule)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve exam schedule from context"))
		return
	}

	var input struct {
		LessonID      int64  `json:"lesson_id" validate:"required,gt=0"`
		TitleOverride string `json:"title_override"`
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

	esi := &store.ExamScopeItem{
		ExamID:        exam.ID,
		LessonID:      input.LessonID,
		TitleOverride: input.TitleOverride,
	}

	err = app.store.ExamScopeItems.Insert(r.Context(), esi)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"exam_scope_item": esi}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listExamScopeItemsHandler(w http.ResponseWriter, r *http.Request) {

	exam, ok := r.Context().Value(examScheduleContextKey).(*store.ExamSchedule)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve exam schedule from context"))
		return
	}

	items, err := app.store.ExamScopeItems.GetAllForExam(r.Context(), exam.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"exam_scope_items": items}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
