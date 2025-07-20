// cmd/api/books.go
package main

import (
	"net/http"
	"strconv"
)

func (app *application) listBooksForCurriculumHandler(w http.ResponseWriter, r *http.Request) {

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

	// Call the data layer to get the books.
	books, err := app.store.Books.GetAllForCurriculum(r.Context(), gradeID, majorID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
