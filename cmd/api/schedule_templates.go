// cmd/api/schedule_templates.go
package main

import (
	"net/http"
	"strconv"
)

func (app *application) listScheduleTemplatesHandler(w http.ResponseWriter, r *http.Request) {

	gradeID, err := strconv.ParseInt(r.URL.Query().Get("grade"), 10, 64)
	if err != nil || gradeID < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	majorID, err := strconv.ParseInt(r.URL.Query().Get("major"), 10, 64)
	if err != nil || majorID < 1 {
		app.badRequestResponse(w, r, err)
		return
	}

	templates, err := app.store.ScheduleTemplates.GetAll(r.Context(), gradeID, majorID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"schedule_templates": templates}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
