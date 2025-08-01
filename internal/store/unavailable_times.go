package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UnavailableTime struct {
	ID          int64
	StudentID   int64
	Title       string
	DayOfWeek   int
	StartTime   time.Time
	EndTime     time.Time
	IsRecurring bool
}

type UnavailableTimeModel struct {
	DB *sql.DB
}

func (m *UnavailableTimeModel) Insert(ctx context.Context, ut *UnavailableTime) error {
	query := `
        INSERT INTO unavailable_times (student_id, title, day_of_week, start_time, end_time, is_recurring)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id
    `
	args := []interface{}{ut.StudentID, ut.Title, ut.DayOfWeek, ut.StartTime, ut.EndTime, ut.IsRecurring}
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ut.ID)
}

func (m *UnavailableTimeModel) Get(ctx context.Context, id int64) (*UnavailableTime, error) {
	query := `
        SELECT id, student_id, title, day_of_week, start_time, end_time, is_recurring
        FROM unavailable_times
        WHERE id = $1
    `
	var ut UnavailableTime
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&ut.ID,
		&ut.StudentID,
		&ut.Title,
		&ut.DayOfWeek,
		&ut.StartTime,
		&ut.EndTime,
		&ut.IsRecurring,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to get unavailable time: %w", err)
	}
	return &ut, nil
}

func (m *UnavailableTimeModel) Update(ctx context.Context, ut *UnavailableTime) error {
	query := `
        UPDATE unavailable_times
        SET title = $1, day_of_week = $2, start_time = $3, end_time = $4, is_recurring = $5
        WHERE id = $6
    `
	args := []interface{}{ut.Title, ut.DayOfWeek, ut.StartTime, ut.EndTime, ut.IsRecurring, ut.ID}
	_, err := m.DB.ExecContext(ctx, query, args...)
	return err
}

func (m *UnavailableTimeModel) Delete(ctx context.Context, id int64) error {
	query := `
        DELETE FROM unavailable_times
        WHERE id = $1
    `
	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}

func (m *UnavailableTimeModel) GetAllForStudent(ctx context.Context, studentID int64) ([]*UnavailableTime, error) {
	query := `
        SELECT id, student_id, title, day_of_week, start_time, end_time, is_recurring
        FROM unavailable_times
        WHERE student_id = $1
        ORDER BY day_of_week, start_time
    `
	rows, err := m.DB.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all unavailable times for student: %w", err)
	}
	defer rows.Close()

	var times []*UnavailableTime
	for rows.Next() {
		var ut UnavailableTime
		err := rows.Scan(
			&ut.ID,
			&ut.StudentID,
			&ut.Title,
			&ut.DayOfWeek,
			&ut.StartTime,
			&ut.EndTime,
			&ut.IsRecurring,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan unavailable time row: %w", err)
		}
		times = append(times, &ut)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return times, nil
}
