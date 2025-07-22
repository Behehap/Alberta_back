package main

import (
	"net/http"
)

func (app *application) listMajorsHandler(w http.ResponseWriter, r *http.Request) {

	majors, err := app.store.Majors.GetAll(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"majors": majors}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
