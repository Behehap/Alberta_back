package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Book struct {
	ID                   int64  `json:"id"`
	Title                string `json:"title"`
	InherentGradeLevelID int64  `json:"inherent_grade_level_id"`
}

type BookModel struct {
	DB *sql.DB
}

func (m *BookModel) Get(ctx context.Context, id int64) (*Book, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `SELECT id, title, inherent_grade_level_id FROM books WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var book Book
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&book.ID,
		&book.Title,
		&book.InherentGradeLevelID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &book, nil
}

// GetAllForCurriculum gets all books for a specific grade and major.
// This uses the book_roles table to figure out the right curriculum.
func (m *BookModel) GetAllForCurriculum(ctx context.Context, gradeID, majorID int64) ([]*Book, error) {
	query := `
        SELECT b.id, b.title, b.inherent_grade_level_id
        FROM books b
        INNER JOIN book_roles br ON b.id = br.book_id
        WHERE br.target_student_grade_id = $1 AND br.major_id = $2
        ORDER BY b.title`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, gradeID, majorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.InherentGradeLevelID,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
