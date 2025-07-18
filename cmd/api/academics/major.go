// cmd/api/majors.go
package main

import (
	"net/http"
)

// listMajorsHandler is our handler for the GET /v1/majors endpoint.
func (app *application) listMajorsHandler(w http.ResponseWriter, r *http.Request) {
	// Call the GetAll method on our Majors store.
	majors, err := app.store.Majors.GetAll(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Wrap the data in an envelope and send it as a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"majors": majors}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
