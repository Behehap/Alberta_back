// cmd/api/students.go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

// createStudentHandler handles the creation of a new student.
func (app *application) createStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Define an anonymous struct to hold the expected JSON input.
	// This helps decouple our API input from our internal data models.
	var input struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number"`
		GradeID     int64  `json:"grade_id" validate:"required,gt=0"`
		MajorID     int64  `json:"major_id" validate:"required,gt=0"`
	}

	// Read and decode the JSON request body.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Validate the input struct.
	// The validator uses the struct tags we defined.
	err = Validate.Struct(input)
	if err != nil {
		// If validation fails, we send a helpful error response.
		// Note: In a real app, you would format these validation errors nicely.
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Copy the data from the input struct to a new Student model.
	student := &store.Student{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		GradeID:     input.GradeID,
		MajorID:     input.MajorID,
	}

	// Insert the new student record into the database.
	err = app.store.Students.Insert(r.Context(), student)
	if err != nil {
		// Handle specific errors, like a duplicate email.
		if errors.Is(err, store.ErrorDuplicateEmail) {
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send a success response to the client.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/students/%d", student.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"student": student}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
