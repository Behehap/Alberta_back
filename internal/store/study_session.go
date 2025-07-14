package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type StudySession struct {
	ID          int64  `json:"id"`
	DailyPlanID int64  `json:"daily_plan_id"`
	LessonID    int64  `json:"lesson_id"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

type StudySessionModel struct {
	DB *sql.DB
}

func (m *StudySessionModel) Insert(ctx context.Context, ss *StudySession) error {
	query := `
        INSERT INTO study_sessions (daily_plan_id, lesson_id, start_time, end_time)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	args := []any{ss.DailyPlanID, ss.LessonID, ss.StartTime, ss.EndTime}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ss.ID)
}

func (m *StudySessionModel) GetAllForDailyPlan(ctx context.Context, dailyPlanID int64) ([]*StudySession, error) {
	query := `
        SELECT id, daily_plan_id, lesson_id, start_time, end_time
        FROM study_sessions
        WHERE daily_plan_id = $1
        ORDER BY start_time ASC`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, dailyPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*StudySession
	for rows.Next() {
		var ss StudySession
		err := rows.Scan(
			&ss.ID,
			&ss.DailyPlanID,
			&ss.LessonID,
			&ss.StartTime,
			&ss.EndTime,
		)
		if err != nil {
			return nil, err
		}
		sessions = append(sessions, &ss)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (m *StudySessionModel) Get(ctx context.Context, id int64) (*StudySession, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, daily_plan_id, lesson_id, start_time, end_time
        FROM study_sessions
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var ss StudySession
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&ss.ID,
		&ss.DailyPlanID,
		&ss.LessonID,
		&ss.StartTime,
		&ss.EndTime,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &ss, nil
}

func (m *StudySessionModel) Update(ctx context.Context, ss *StudySession) error {
	query := `
        UPDATE study_sessions
        SET lesson_id = $1, start_time = $2, end_time = $3
        WHERE id = $4`

	args := []any{ss.LessonID, ss.StartTime, ss.EndTime, ss.ID}

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

func (m *StudySessionModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM study_sessions WHERE id = $1`

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
