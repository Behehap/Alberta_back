package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createExamScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title         string `json:"title" validate:"required"`
		ExamDate      string `json:"exam_date" validate:"required"` // Expects "YYYY-MM-DD"
		Organisation  string `json:"organisation"`
		TargetGradeID int64  `json:"target_grade_id" validate:"required,gt=0"`
		MajorID       int64  `json:"major_id" validate:"required,gt=0"`
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

	examDate, err := time.Parse("2006-01-02", input.ExamDate)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid date format for exam_date, please use YYYY-MM-DD"))
		return
	}

	es := &store.ExamSchedule{
		Title:         input.Title,
		ExamDate:      examDate,
		Organisation:  input.Organisation,
		TargetGradeID: input.TargetGradeID,
		MajorID:       input.MajorID,
	}

	err = app.store.ExamSchedules.Insert(r.Context(), es)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"exam_schedule": es}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listExamSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	exams, err := app.store.ExamSchedules.GetAllForStudentCurriculum(r.Context(), student.GradeID, student.MajorID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"exam_schedules": exams}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getExamScheduleHandler(w http.ResponseWriter, r *http.Request) {
	exam, ok := r.Context().Value(examScheduleContextKey).(*store.ExamSchedule)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve exam schedule from context"))
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"exam_schedule": exam}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
