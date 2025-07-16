package main

import (
	"errors"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createSessionReportHandler(w http.ResponseWriter, r *http.Request) {
	studyItem, ok := r.Context().Value(weeklyStudyItemContextKey).(*store.WeeklyStudyItem)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly study item from context"))
		return
	}

	var input struct {
		IsReview      bool    `json:"is_review"`
		NumTests      int     `json:"num_tests"`
		NumWrongTests int     `json:"num_wrong_tests"`
		SessionScore  float64 `json:"session_score"`
		Notes         string  `json:"notes"`
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

	sr := &store.SessionReport{
		WeeklyStudyItemID: studyItem.ID,
		IsReview:          input.IsReview,
		NumTests:          input.NumTests,
		NumWrongTests:     input.NumWrongTests,
		SessionScore:      input.SessionScore,
		Notes:             input.Notes,
	}

	err = app.store.SessionReports.Insert(r.Context(), sr)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"session_report": sr}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSessionReportHandler(w http.ResponseWriter, r *http.Request) {
	studyItem, ok := r.Context().Value(weeklyStudyItemContextKey).(*store.WeeklyStudyItem)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve weekly study item from context"))
		return
	}

	report, err := app.store.SessionReports.GetForWeeklyStudyItem(r.Context(), studyItem.ID)
	if err != nil {
		if errors.Is(err, store.ErrorNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"session_report": report}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
