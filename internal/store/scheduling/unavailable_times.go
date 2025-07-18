// internal/store/unavailable_times.go
package store

import (
	"context"
	"database/sql"
	"time"
)

// UnavailableTime is a block of time when a student is busy.
type UnavailableTime struct {
	ID          int64  `json:"id"`
	StudentID   int64  `json:"student_id"`
	Title       string `json:"title,omitempty"`
	DayOfWeek   int    `json:"day_of_week"` // Using Go's time.Weekday: Sunday=0, Monday=1, etc.
	StartTime   string `json:"start_time"`  // Storing as "HH:MM:SS".
	EndTime     string `json:"end_time"`
	IsRecurring bool   `json:"is_recurring"`
}

// UnavailableTimeModel holds the database connection.
type UnavailableTimeModel struct {
	DB *sql.DB
}

// Insert adds a new unavailable time for a student.
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

// GetAllForStudent gets all unavailable time slots for a specific student.
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
			return nil, err
		}
		times = append(times, &ut)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return times, nil
}

// Update modifies an existing unavailable time slot.
func (m *UnavailableTimeModel) Update(ctx context.Context, ut *UnavailableTime) error {
	query := `
        UPDATE unavailable_times
        SET title = $1, day_of_week = $2, start_time = $3, end_time = $4, is_recurring = $5
        WHERE id = $6 AND student_id = $7` // Ensure student can only update their own times

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

// Delete removes an unavailable time slot.
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
