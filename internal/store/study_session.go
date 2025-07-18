package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type StudySession struct {
	ID             int64     `json:"id"`
	DailyPlanID    int64     `json:"daily_plan_id"`
	BookID         int64     `json:"book_id"`
	IsCompleted    bool      `json:"is_completed"`
	CompletionDate time.Time `json:"completion_date,omitempty"` // Nullable in DB
	StartTime      string    `json:"start_time"`                // Storing as "HH:MM:SS"
	EndTime        string    `json:"end_time"`                  // Storing as "HH:MM:SS"
}

type StudySessionModel struct {
	DB *sql.DB
}

func (m *StudySessionModel) Insert(ctx context.Context, ss *StudySession) error {
	query := `
        INSERT INTO study_sessions (daily_plan_id, book_id, start_time, end_time, is_completed)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, is_completed`

	args := []any{ss.DailyPlanID, ss.BookID, ss.StartTime, ss.EndTime, ss.IsCompleted}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ss.ID, &ss.IsCompleted)
}

func (m *StudySessionModel) Get(ctx context.Context, id int64) (*StudySession, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}
	query := `
        SELECT id, daily_plan_id, book_id, is_completed, completion_date, start_time, end_time
        FROM study_sessions
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var ss StudySession
	var completionDate sql.NullTime
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&ss.ID,
		&ss.DailyPlanID,
		&ss.BookID,
		&ss.IsCompleted,
		&completionDate,
		&ss.StartTime,
		&ss.EndTime,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	if completionDate.Valid {
		ss.CompletionDate = completionDate.Time
	}

	return &ss, nil
}

func (m *StudySessionModel) GetAllForDailyPlan(ctx context.Context, dailyPlanID int64) ([]*StudySession, error) {
	query := `
        SELECT id, daily_plan_id, book_id, is_completed, completion_date, start_time, end_time
        FROM study_sessions
        WHERE daily_plan_id = $1
        ORDER BY start_time`

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
		var completionDate sql.NullTime
		err := rows.Scan(
			&ss.ID,
			&ss.DailyPlanID,
			&ss.BookID,
			&ss.IsCompleted,
			&completionDate,
			&ss.StartTime,
			&ss.EndTime,
		)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			ss.CompletionDate = completionDate.Time
		}
		sessions = append(sessions, &ss)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return sessions, nil
}

func (m *StudySessionModel) Update(ctx context.Context, ss *StudySession) error {
	query := `
        UPDATE study_sessions
        SET daily_plan_id = $1, book_id = $2, is_completed = $3, completion_date = $4, start_time = $5, end_time = $6
        WHERE id = $7`

	var completionDate sql.NullTime
	if ss.IsCompleted {
		completionDate.Time = time.Now()
		completionDate.Valid = true
	}

	args := []any{ss.DailyPlanID, ss.BookID, ss.IsCompleted, completionDate, ss.StartTime, ss.EndTime, ss.ID}
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

// Delete removes a study session.
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
