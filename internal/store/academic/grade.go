// internal/store/grades.go
package store

import (
	"context"
	"database/sql"
	"time"
)

// Grade represents a single academic grade level.
type Grade struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// GradeModel holds the database connection.
type GradeModel struct {
	DB *sql.DB
}

// Get retrieves a single grade from the database by its ID.
func (m *GradeModel) Get(ctx context.Context, id int64) (*Grade, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `SELECT id, name FROM grades WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var grade Grade
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&grade.ID, &grade.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &grade, nil
}

// GetAll returns a slice of all grades, ordered by their ID.
func (m *GradeModel) GetAll(ctx context.Context) ([]*Grade, error) {
	query := `SELECT id, name FROM grades ORDER BY id`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create an empty slice to hold the grades data.
	var grades []*Grade

	// Iterate over the rows, scanning the data into a Grade struct.
	for rows.Next() {
		var grade Grade
		if err := rows.Scan(&grade.ID, &grade.Name); err != nil {
			return nil, err
		}
		grades = append(grades, &grade)
	}

	// When the rows.Next() loop has finished, call rows.Err() to check if any
	// error occurred during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return grades, nil
}
