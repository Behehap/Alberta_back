package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) listLessonsForBookHandler(w http.ResponseWriter, r *http.Request) {
	// Get the book ID from the URL, e.g. /v1/books/1/lessons
	bookID, err := strconv.ParseInt(chi.URLParam(r, "bookID"), 10, 64)
	if err != nil || bookID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// Call the data layer to get the lessons for this book.
	lessons, err := app.store.Lessons.GetAllForBook(r.Context(), bookID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send the list of lessons back to the client.
	err = app.writeJSON(w, http.StatusOK, envelope{"lessons": lessons}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
