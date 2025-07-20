package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Lesson struct {
	ID                        int64         `json:"id"`
	Name                      string        `json:"name"`
	BookID                    int64         `json:"book_id"`
	EstimatedStudyTimeMinutes sql.NullInt64 `json:"estimated_study_time_minutes,omitempty"`
}

type LessonModel struct {
	DB *sql.DB
}

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

func (m *LessonModel) GetAllForBook(ctx context.Context, bookID int64) ([]*Lesson, error) {
	query := `
        SELECT id, name, book_id, estimated_study_time_minutes
        FROM lessons
        WHERE book_id = $1
        ORDER BY id`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []*Lesson
	for rows.Next() {
		var lesson Lesson
		err := rows.Scan(
			&lesson.ID,
			&lesson.Name,
			&lesson.BookID,
			&lesson.EstimatedStudyTimeMinutes,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}
