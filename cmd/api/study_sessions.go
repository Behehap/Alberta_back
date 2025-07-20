package main

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) createStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	dailyPlanID, err := strconv.ParseInt(chi.URLParam(r, "dailyPlanID"), 10, 64)
	if err != nil || dailyPlanID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		BookID    int64  `json:"book_id" validate:"required,gt=0"`
		StartTime string `json:"start_time" validate:"required"`
		EndTime   string `json:"end_time" validate:"required"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(input)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	parsedStartTime, err := time.Parse("15:04", input.StartTime)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid start_time format, please use HH:MM"))
		return
	}
	parsedEndTime, err := time.Parse("15:04", input.EndTime)
	if err != nil {
		app.badRequestResponse(w, r, errors.New("invalid end_time format, please use HH:MM"))
		return
	}

	ss := &store.StudySession{
		DailyPlanID:    dailyPlanID,
		BookID:         input.BookID,
		StartTime:      parsedStartTime.Format("15:04:05"),
		EndTime:        parsedEndTime.Format("15:04:05"),
		IsCompleted:    false,
		CompletionDate: sql.NullTime{}, // Initialize as null
	}

	err = app.store.StudySessions.Insert(r.Context(), ss)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"study_session": ss}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listStudySessionsHandler(w http.ResponseWriter, r *http.Request) {
	dailyPlanID, err := strconv.ParseInt(chi.URLParam(r, "dailyPlanID"), 10, 64)
	if err != nil || dailyPlanID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	sessions, err := app.store.StudySessions.GetAllForDailyPlan(r.Context(), dailyPlanID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"study_sessions": sessions}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	studySession, ok := r.Context().Value(studySessionContextKey).(*store.StudySession)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve study session from context"))
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"study_session": studySession}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(chi.URLParam(r, "sessionID"), 10, 64)
	if err != nil || sessionID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	session, err := app.store.StudySessions.Get(r.Context(), sessionID)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		IsCompleted *bool   `json:"is_completed"`
		BookID      *int64  `json:"book_id"`
		StartTime   *string `json:"start_time"`
		EndTime     *string `json:"end_time"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.IsCompleted != nil {
		session.IsCompleted = *input.IsCompleted
		if *input.IsCompleted {
			session.CompletionDate = sql.NullTime{Time: time.Now(), Valid: true}
		} else {
			session.CompletionDate = sql.NullTime{Valid: false}
		}
	}
	if input.BookID != nil {
		session.BookID = *input.BookID
	}
	if input.StartTime != nil {
		parsedTime, err := time.Parse("15:04", *input.StartTime)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("invalid start_time format, please use HH:MM"))
			return
		}
		session.StartTime = parsedTime.Format("15:04:05")
	}
	if input.EndTime != nil {
		parsedTime, err := time.Parse("15:04", *input.EndTime)
		if err != nil {
			app.badRequestResponse(w, r, errors.New("invalid end_time format, please use HH:MM"))
			return
		}
		session.EndTime = parsedTime.Format("15:04:05")
	}

	err = app.store.StudySessions.Update(r.Context(), session)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"study_session": session}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteStudySessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := strconv.ParseInt(chi.URLParam(r, "sessionID"), 10, 64)
	if err != nil || sessionID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.store.StudySessions.Delete(r.Context(), sessionID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "study session successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
