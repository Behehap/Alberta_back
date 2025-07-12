// internal/store/books.go
package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// Book represents a single textbook.
type Book struct {
	ID                   int64  `json:"id"`
	Title                string `json:"title"`
	InherentGradeLevelID int64  `json:"inherent_grade_level_id"`
}

// BookModel holds the database connection and implements the BookStore interface.
type BookModel struct {
	DB *sql.DB
}

// Get retrieves a single book from the database by its ID.
func (m *BookModel) Get(ctx context.Context, id int64) (*Book, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, title, inherent_grade_level_id
        FROM books
        WHERE id = $1`

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
