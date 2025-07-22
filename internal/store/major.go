package store

import (
	"context"
	"database/sql"
	"time"
)

type Major struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MajorModel struct {
	DB *sql.DB
}

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

func (m *MajorModel) GetAll(ctx context.Context) ([]*Major, error) {
	query := `SELECT id, name FROM majors ORDER BY id`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var majors []*Major
	for rows.Next() {
		var major Major
		if err := rows.Scan(&major.ID, &major.Name); err != nil {
			return nil, err
		}
		majors = append(majors, &major)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return majors, nil
}
