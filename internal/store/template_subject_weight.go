package store

import (
	"context"
	"database/sql"
	"time"
)

type TemplateSubjectWeight struct {
	ID         int64   `json:"id"`
	TemplateID int64   `json:"template_id"`
	BookID     int64   `json:"book_id"`
	Weight     float64 `json:"weight"`
}

type TemplateSubjectWeightModel struct {
	DB *sql.DB
}

func (m *TemplateSubjectWeightModel) GetWeightsForTemplate(ctx context.Context, templateID int64) ([]*TemplateSubjectWeight, error) {
	query := `
        SELECT id, template_id, book_id, weight
        FROM template_subject_weights
        WHERE template_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var weights []*TemplateSubjectWeight
	for rows.Next() {
		var weight TemplateSubjectWeight
		err := rows.Scan(
			&weight.ID,
			&weight.TemplateID,
			&weight.BookID,
			&weight.Weight,
		)
		if err != nil {
			return nil, err
		}
		weights = append(weights, &weight)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return weights, nil
}

func (m *TemplateSubjectWeightModel) SetWeight(ctx context.Context, weight *TemplateSubjectWeight) error {
	query := `
        INSERT INTO template_subject_weights (template_id, book_id, weight)
        VALUES ($1, $2, $3)
        ON CONFLICT (template_id, book_id) 
        DO UPDATE SET weight = $3
        RETURNING id`

	args := []any{weight.TemplateID, weight.BookID, weight.Weight}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&weight.ID)
}
