// cmd/api/grades.go
package main

import (
	"net/http"
)

func (app *application) listGradesHandler(w http.ResponseWriter, r *http.Request) {
	grades, err := app.store.Grades.GetAll(r.Context())
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"grades": grades}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
