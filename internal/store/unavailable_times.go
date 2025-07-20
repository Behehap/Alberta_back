package store

import (
	"context"
	"database/sql"
	"time"
)

type UnavailableTime struct {
	ID          int64  `json:"id"`
	StudentID   int64  `json:"student_id"`
	Title       string `json:"title,omitempty"`
	DayOfWeek   int    `json:"day_of_week"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	IsRecurring bool   `json:"is_recurring"`
}

type UnavailableTimeModel struct {
	DB *sql.DB
}

func (m *UnavailableTimeModel) Insert(ctx context.Context, ut *UnavailableTime) error {
	query := `
        INSERT INTO unavailable_times (student_id, title, day_of_week, start_time, end_time, is_recurring)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	args := []any{ut.StudentID, ut.Title, ut.DayOfWeek, ut.StartTime, ut.EndTime, ut.IsRecurring}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&ut.ID)
}

func (m *UnavailableTimeModel) GetAllForStudent(ctx context.Context, studentID int64) ([]*UnavailableTime, error) {
	query := `
        SELECT id, student_id, title, day_of_week, start_time, end_time, is_recurring
        FROM unavailable_times
        WHERE student_id = $1
        ORDER BY day_of_week, start_time`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var times []*UnavailableTime
	for rows.Next() {
		var ut UnavailableTime
		var dbStartTime, dbEndTime time.Time
		err := rows.Scan(
			&ut.ID,
			&ut.StudentID,
			&ut.Title,
			&ut.DayOfWeek,
			&dbStartTime,
			&dbEndTime,
			&ut.IsRecurring,
		)
		if err != nil {
			return nil, err
		}

		ut.StartTime = dbStartTime.Format("15:04:05")
		ut.EndTime = dbEndTime.Format("15:04:05")

		times = append(times, &ut)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return times, nil
}

func (m *UnavailableTimeModel) Update(ctx context.Context, ut *UnavailableTime) error {
	query := `
        UPDATE unavailable_times
        SET title = $1, day_of_week = $2, start_time = $3, end_time = $4, is_recurring = $5
        WHERE id = $6 AND student_id = $7`

	args := []any{ut.Title, ut.DayOfWeek, ut.StartTime, ut.EndTime, ut.IsRecurring, ut.ID, ut.StudentID}

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

func (m *UnavailableTimeModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM unavailable_times WHERE id = $1`

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
