// cmd/api/students.go
package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Behehap/Alberta/internal/store"
)

func (app *application) createStudentHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		FirstName   string `json:"first_name" validate:"required"`
		LastName    string `json:"last_name" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number"`
		GradeID     int64  `json:"grade_id" validate:"required,gt=0"`
		MajorID     int64  `json:"major_id" validate:"required,gt=0"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = Validate.Struct(input)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	student := &store.Student{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		GradeID:     input.GradeID,
		MajorID:     input.MajorID,
	}

	err = app.store.Students.Insert(r.Context(), student)
	if err != nil {

		if errors.Is(err, store.ErrorDuplicateEmail) {
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/students/%d", student.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"student": student}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getStudentHandler(w http.ResponseWriter, r *http.Request) {

	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateStudentHandler(w http.ResponseWriter, r *http.Request) {

	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	var input struct {
		FirstName   *string `json:"first_name"`
		LastName    *string `json:"last_name"`
		Email       *string `json:"email" validate:"omitempty,email"`
		PhoneNumber *string `json:"phone_number"`
		GradeID     *int64  `json:"grade_id" validate:"omitempty,gt=0"`
		MajorID     *int64  `json:"major_id" validate:"omitempty,gt=0"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

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

	err = Validate.Struct(student)
	if err != nil {
		app.failedValidationResponse(w, r, map[string]string{"error": err.Error()})
		return
	}

	err = app.store.Students.Update(r.Context(), student)
	if err != nil {
		if errors.Is(err, store.ErrorDuplicateEmail) {
			app.failedValidationResponse(w, r, map[string]string{"email": "a user with this email address already exists"})
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteStudentHandler(w http.ResponseWriter, r *http.Request) {

	student, ok := r.Context().Value(studentContextKey).(*store.Student)
	if !ok {
		app.serverErrorResponse(w, r, errors.New("could not retrieve student from context"))
		return
	}

	err := app.store.Students.Delete(r.Context(), student.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "student successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
