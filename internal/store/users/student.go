// internal/store/students.go
package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Student represents a single student user.
type Student struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`
	GradeID     int64  `json:"grade_id"`
	MajorID     int64  `json:"major_id"`
}

// StudentModel holds the database connection and implements the StudentStore interface.
type StudentModel struct {
	DB *sql.DB
}

// Insert adds a new student record to the database.
func (m *StudentModel) Insert(ctx context.Context, student *Student) error {
	query := `
        INSERT INTO students (first_name, last_name, email, phone_number, grade_id, major_id)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	args := []any{student.FirstName, student.LastName, student.Email, student.PhoneNumber, student.GradeID, student.MajorID}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	// Use QueryRowContext to execute the query and scan the returned ID.
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&student.ID)
	if err != nil {
		// Check for a specific database constraint violation for duplicate emails.
		// The exact error string might depend on your database driver (e.g., "unique_violation").
		if err.Error() == `pq: duplicate key value violates unique constraint "students_email_key"` {
			return ErrorDuplicateEmail
		}
		return err
	}
	return nil
}

// Get retrieves a single student from the database by their ID.
func (m *StudentModel) Get(ctx context.Context, id int64) (*Student, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, first_name, last_name, email, phone_number, grade_id, major_id
        FROM students
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var s Student
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&s.ID,
		&s.FirstName,
		&s.LastName,
		&s.Email,
		&s.PhoneNumber,
		&s.GradeID,
		&s.MajorID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &s, nil
}

// Update modifies an existing student record.
func (m *StudentModel) Update(ctx context.Context, student *Student) error {
	query := `
        UPDATE students
        SET first_name = $1, last_name = $2, email = $3, phone_number = $4, grade_id = $5, major_id = $6
        WHERE id = $7`

	args := []any{
		student.FirstName,
		student.LastName,
		student.Email,
		student.PhoneNumber,
		student.GradeID,
		student.MajorID,
		student.ID,
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "students_email_key"` {
			return ErrorDuplicateEmail
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}

// Delete removes a student record from the database by their ID.
func (m *StudentModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM students WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNotFound
	}

	return nil
}
