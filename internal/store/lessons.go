// internal/store/lessons.go
package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Lesson represents a single lesson or chapter within a book.
type Lesson struct {
	ID                        int64  `json:"id"`
	Name                      string `json:"name"`
	BookID                    int64  `json:"book_id"`
	EstimatedStudyTimeMinutes int    `json:"estimated_study_time_minutes,omitempty"`
}

// LessonModel holds the database connection and implements the LessonStore interface.
type LessonModel struct {
	DB *sql.DB
}

// Get retrieves a single lesson from the database by its ID.
func (m *LessonModel) Get(ctx context.Context, id int64) (*Lesson, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, name, book_id, estimated_study_time_minutes
        FROM lessons
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var lesson Lesson
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&lesson.ID,
		&lesson.Name,
		&lesson.BookID,
		&lesson.EstimatedStudyTimeMinutes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &lesson, nil
}
