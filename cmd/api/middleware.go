// cmd/api/middleware.go
package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Behehap/Alberta/internal/store"

	"github.com/go-chi/chi/v5"
)

// studentContextMiddleware is a helper that runs before our main handlers.
// Its job is to grab the student ID from the URL, fetch the student
// from the database, and then stick that student object into the request's
// context. This way, our other handlers don't have to repeat this logic.
func (app *application) studentContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// chi.URLParam gets the {studentID} part from the URL.
		studentID, err := strconv.ParseInt(chi.URLParam(r, "studentID"), 10, 64)
		if err != nil || studentID < 1 {
			app.notFoundResponse(w, r)
			return
		}

		// Use our store to fetch the student from the database.
		student, err := app.store.Students.Get(r.Context(), studentID)
		if err != nil {
			// If we get an error (like the student not existing),
			// we just send a 404 Not Found response.
			if err == store.ErrorNotFound {
				app.notFoundResponse(w, r)
				return
			}
			// For any other kind of error, it's a server problem.
			app.serverErrorResponse(w, r, err)
			return
		}

		// Add the student to the request's context using our custom key.
		ctx := context.WithValue(r.Context(), studentContextKey, student)

		// Call the next handler in the chain, passing it the new context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
