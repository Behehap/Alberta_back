package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type WeeklyPlan struct {
	ID                       int64        `json:"id"`
	StudentID                int64        `json:"student_id"`
	StartDateOfWeek          time.Time    `json:"start_date_of_week"`
	DayStartTime             sql.NullTime `json:"day_start_time,omitempty"`
	MaxStudyTimeHoursPerWeek int          `json:"max_study_time_hours_per_week,omitempty"`
}

type WeeklyPlanModel struct {
	DB *sql.DB
}

func (m *WeeklyPlanModel) Insert(ctx context.Context, wp *WeeklyPlan) error {
	query := `
        INSERT INTO weekly_plans (student_id, start_date_of_week, day_start_time, max_study_time_hours_per_week)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	args := []any{wp.StudentID, wp.StartDateOfWeek, wp.DayStartTime, wp.MaxStudyTimeHoursPerWeek}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&wp.ID)
}

func (m *WeeklyPlanModel) Get(ctx context.Context, id int64) (*WeeklyPlan, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, student_id, start_date_of_week, day_start_time, max_study_time_hours_per_week
        FROM weekly_plans
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var wp WeeklyPlan
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&wp.ID,
		&wp.StudentID,
		&wp.StartDateOfWeek,
		&wp.DayStartTime,
		&wp.MaxStudyTimeHoursPerWeek,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &wp, nil
}

func (m *WeeklyPlanModel) GetAllForStudent(ctx context.Context, studentID int64) ([]*WeeklyPlan, error) {
	query := `
        SELECT id, student_id, start_date_of_week, day_start_time, max_study_time_hours_per_week
        FROM weekly_plans
        WHERE student_id = $1
        ORDER BY start_date_of_week DESC`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []*WeeklyPlan
	for rows.Next() {
		var wp WeeklyPlan
		err := rows.Scan(
			&wp.ID,
			&wp.StudentID,
			&wp.StartDateOfWeek,
			&wp.DayStartTime,
			&wp.MaxStudyTimeHoursPerWeek,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, &wp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plans, nil
}

func (m *WeeklyPlanModel) Update(ctx context.Context, wp *WeeklyPlan) error {
	query := `
        UPDATE weekly_plans
        SET start_date_of_week = $1, day_start_time = $2, max_study_time_hours_per_week = $3
        WHERE id = $4 AND student_id = $5`

	args := []any{wp.StartDateOfWeek, wp.DayStartTime, wp.MaxStudyTimeHoursPerWeek, wp.ID, wp.StudentID}

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

func (m *WeeklyPlanModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM weekly_plans WHERE id = $1`

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
