package store

import (
	"context"
	"database/sql"
	"time"
)

type SubjectFrequency struct {
	ID               int64 `json:"id"`
	WeeklyPlanID     int64 `json:"weekly_plan_id"`
	BookID           int64 `json:"book_id"`
	FrequencyPerWeek int   `json:"frequency_per_week"`
}

type SubjectFrequencyModel struct {
	DB *sql.DB
}

func (m *SubjectFrequencyModel) Insert(ctx context.Context, sf *SubjectFrequency) error {
	query := `
        INSERT INTO subject_frequencies (weekly_plan_id, book_id, frequency_per_week)
        VALUES ($1, $2, $3)
        RETURNING id`

	args := []any{sf.WeeklyPlanID, sf.BookID, sf.FrequencyPerWeek}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&sf.ID)
}

func (m *SubjectFrequencyModel) GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*SubjectFrequency, error) {
	query := `
        SELECT id, weekly_plan_id, book_id, frequency_per_week
        FROM subject_frequencies
        WHERE weekly_plan_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, weeklyPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var frequencies []*SubjectFrequency
	for rows.Next() {
		var sf SubjectFrequency
		err := rows.Scan(
			&sf.ID,
			&sf.WeeklyPlanID,
			&sf.BookID,
			&sf.FrequencyPerWeek,
		)
		if err != nil {
			return nil, err
		}
		frequencies = append(frequencies, &sf)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return frequencies, nil
}

func (m *SubjectFrequencyModel) Update(ctx context.Context, sf *SubjectFrequency) error {
	query := `
        UPDATE subject_frequencies
        SET book_id = $1, frequency_per_week = $2
        WHERE id = $3 AND weekly_plan_id = $4`

	args := []any{sf.BookID, sf.FrequencyPerWeek, sf.ID, sf.WeeklyPlanID}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
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

func (m *SubjectFrequencyModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM subject_frequencies WHERE id = $1`

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
