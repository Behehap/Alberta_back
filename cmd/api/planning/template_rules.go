package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/Behehap/Alberta/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *application) createTemplateRuleHandler(w http.ResponseWriter, r *http.Request) {
	templateID, err := strconv.ParseInt(chi.URLParam(r, "templateID"), 10, 64)
	if err != nil || templateID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		BookID              int64   `json:"book_id" validate:"required,gt=0"`
		DefaultFrequency    int     `json:"default_frequency" validate:"required,gt=0"`
		SchedulingHints     *string `json:"scheduling_hints"`
		ConsecutiveSessions *bool   `json:"consecutive_sessions"`
		TimePreference      *string `json:"time_preference"`
		PrioritySlot        *string `json:"priority_slot"`
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

	tr := &store.TemplateRule{
		TemplateID:       templateID,
		BookID:           input.BookID,
		DefaultFrequency: input.DefaultFrequency,
	}

	if input.SchedulingHints != nil {
		tr.SchedulingHints = sql.NullString{String: *input.SchedulingHints, Valid: true}
	}
	if input.ConsecutiveSessions != nil {
		tr.ConsecutiveSessions = sql.NullBool{Bool: *input.ConsecutiveSessions, Valid: true}
	}
	if input.TimePreference != nil {
		tr.TimePreference = sql.NullString{String: strings.ToLower(*input.TimePreference), Valid: true}
	}
	if input.PrioritySlot != nil {
		tr.PrioritySlot = sql.NullString{String: strings.ToLower(*input.PrioritySlot), Valid: true}
	}

	err = app.store.TemplateRules.Insert(r.Context(), tr)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"template_rule": tr}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listTemplateRulesHandler(w http.ResponseWriter, r *http.Request) {
	templateID, err := strconv.ParseInt(chi.URLParam(r, "templateID"), 10, 64)
	if err != nil || templateID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	rules, err := app.store.TemplateRules.GetAllForTemplate(r.Context(), templateID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"template_rules": rules}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateTemplateRuleHandler(w http.ResponseWriter, r *http.Request) {
	templateID, err := strconv.ParseInt(chi.URLParam(r, "templateID"), 10, 64)
	if err != nil || templateID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	ruleID, err := strconv.ParseInt(chi.URLParam(r, "ruleID"), 10, 64)
	if err != nil || ruleID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	rule, err := app.store.TemplateRules.Get(r.Context(), ruleID)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if rule.TemplateID != templateID {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		BookID              *int64  `json:"book_id"`
		DefaultFrequency    *int    `json:"default_frequency"`
		SchedulingHints     *string `json:"scheduling_hints"`
		ConsecutiveSessions *bool   `json:"consecutive_sessions"`
		TimePreference      *string `json:"time_preference"`
		PrioritySlot        *string `json:"priority_slot"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.BookID != nil {
		rule.BookID = *input.BookID
	}
	if input.DefaultFrequency != nil {
		rule.DefaultFrequency = *input.DefaultFrequency
	}
	if input.SchedulingHints != nil {
		rule.SchedulingHints = sql.NullString{String: *input.SchedulingHints, Valid: true}
	} else {
		rule.SchedulingHints = sql.NullString{}
	}
	if input.ConsecutiveSessions != nil {
		rule.ConsecutiveSessions = sql.NullBool{Bool: *input.ConsecutiveSessions, Valid: true}
	} else {
		rule.ConsecutiveSessions = sql.NullBool{}
	}
	if input.TimePreference != nil {
		rule.TimePreference = sql.NullString{String: strings.ToLower(*input.TimePreference), Valid: true}
	} else {
		rule.TimePreference = sql.NullString{}
	}
	if input.PrioritySlot != nil {
		rule.PrioritySlot = sql.NullString{String: strings.ToLower(*input.PrioritySlot), Valid: true}
	} else {
		rule.PrioritySlot = sql.NullString{}
	}

	err = Validate.Struct(rule)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	err = app.store.TemplateRules.Update(r.Context(), rule)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"template_rule": rule}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTemplateRuleHandler(w http.ResponseWriter, r *http.Request) {
	templateID, err := strconv.ParseInt(chi.URLParam(r, "templateID"), 10, 64)
	if err != nil || templateID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	ruleID, err := strconv.ParseInt(chi.URLParam(r, "ruleID"), 10, 64)
	if err != nil || ruleID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	rule, err := app.store.TemplateRules.Get(r.Context(), ruleID)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if rule.TemplateID != templateID {
		app.notFoundResponse(w, r)
		return
	}

	err = app.store.TemplateRules.Delete(r.Context(), ruleID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "template rule successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
