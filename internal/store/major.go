// internal/store/majors.go
package store

import (
	"context"
	"database/sql"
	"time"
)

// Major represents a single academic major.
type Major struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// MajorModel holds the database connection.
type MajorModel struct {
	DB *sql.DB
}

// Get retrieves a single major from the database by its ID.
func (m *MajorModel) Get(ctx context.Context, id int64) (*Major, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `SELECT id, name FROM majors WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var major Major
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&major.ID, &major.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &major, nil
}
