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
	// This struct is just for reading the incoming JSON.
	var input struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number"`
		GradeID     int64  `json:"grade_id" validate:"required,gt=0"`
		MajorID     int64  `json:"major_id" validate:"required,gt=0"`
	}

	// Try to read the JSON from the request.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Check the input data to make sure it's valid.
	err = Validate.Struct(input)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Copy the data over to our actual Student model.
	student := &store.Student{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		GradeID:     input.GradeID,
		MajorID:     input.MajorID,
	}

	// Try to save the new student to the database.
	err = app.store.Students.Insert(r.Context(), student)
	if err != nil {
		// Check if it's a duplicate email error.
		if errors.Is(err, store.ErrorDuplicateEmail) {
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send back a 201 Created response with the new student's data.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/students/%d", student.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"student": student}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// getStudentHandler just grabs the student we put in the context
// with our middleware and sends it back to the client.
func (app *application) getStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the student from the request context.
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	// Write the student data as a JSON response.
	err := app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// updateStudentHandler handles partial updates for a student.
func (app *application) updateStudentHandler(w http.ResponseWriter, r *http.Request) {
	// First, get the existing student record from the context.
	// Our middleware already handled fetching it and the 404 case.
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	// This input struct uses pointers so we can tell the difference
	// between a field that's missing and a field that's deliberately set to "".
	var input struct {
		FirstName   *string `json:"first_name"`
		LastName    *string `json:"last_name"`
		Email       *string `json:"email" validate:"omitempty,email"`
		PhoneNumber *string `json:"phone_number"`
		GradeID     *int64  `json:"grade_id" validate:"omitempty,gt=0"`
		MajorID     *int64  `json:"major_id" validate:"omitempty,gt=0"`
	}

	// Read the JSON from the request.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Check each field in the input. If it's not nil, the user sent
	// a new value, so we update the student record we fetched earlier.
	if input.FirstName != nil {
		student.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		student.LastName = *input.LastName
	}
	if input.Email != nil {
		student.Email = *input.Email
	}
	if input.PhoneNumber != nil {
		student.PhoneNumber = *input.PhoneNumber
	}
	if input.GradeID != nil {
		student.GradeID = *input.GradeID
	}
	if input.MajorID != nil {
		student.MajorID = *input.MajorID
	}

	// Validate the updated student record.
	err = Validate.Struct(student)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Save the updated record back to the database.
	err = app.store.Students.Update(r.Context(), student)
	if err != nil {
		if errors.Is(err, store.ErrorDuplicateEmail) {
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send the updated student back to the client.
	err = app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// deleteStudentHandler removes a student from the database.
func (app *application) deleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Get the student from the request context.
	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	// Call the store to delete the student.
	err := app.store.Students.Delete(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Send back a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "student successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
