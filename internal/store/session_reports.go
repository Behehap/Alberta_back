// internal/store/session_reports.go
package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// SessionReport contains the details of a completed study session.
type SessionReport struct {
	ID            int64   `json:"id"`
	StudyBlockID  int64   `json:"study_block_id"`
	IsCompleted   bool    `json:"is_completed"`
	IsReview      bool    `json:"is_review"`
	NumTests      int     `json:"num_tests,omitempty"`
	NumWrongTests int     `json:"num_wrong_tests,omitempty"`
	SessionScore  float64 `json:"session_score,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}

// SessionReportModel holds the database connection.
type SessionReportModel struct {
	DB *sql.DB
}

// Insert creates a new session report for a specific study block.
func (m *SessionReportModel) Insert(ctx context.Context, sr *SessionReport) error {
	query := `
        INSERT INTO session_reports (study_block_id, is_completed, is_review, num_tests, num_wrong_tests, session_score, notes)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	args := []any{sr.StudyBlockID, sr.IsCompleted, sr.IsReview, sr.NumTests, sr.NumWrongTests, sr.SessionScore, sr.Notes}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&sr.ID)
}

// GetForStudyBlock retrieves the report for a specific study block.
// Since there's a one-to-one relationship, we expect at most one report.
func (m *SessionReportModel) GetForStudyBlock(ctx context.Context, studyBlockID int64) (*SessionReport, error) {
	if studyBlockID < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, study_block_id, is_completed, is_review, num_tests, num_wrong_tests, session_score, notes
        FROM session_reports
        WHERE study_block_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var sr SessionReport
	err := m.DB.QueryRowContext(ctx, query, studyBlockID).Scan(
		&sr.ID,
		&sr.StudyBlockID,
		&sr.IsCompleted,
		&sr.IsReview,
		&sr.NumTests,
		&sr.NumWrongTests,
		&sr.SessionScore,
		&sr.Notes,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &sr, nil
}

// Update modifies an existing session report.
func (m *SessionReportModel) Update(ctx context.Context, sr *SessionReport) error {
	query := `
        UPDATE session_reports
        SET is_completed = $1, is_review = $2, num_tests = $3, num_wrong_tests = $4, session_score = $5, notes = $6
        WHERE id = $7`

	args := []any{sr.IsCompleted, sr.IsReview, sr.NumTests, sr.NumWrongTests, sr.SessionScore, sr.Notes, sr.ID}

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

// Delete removes a session report.
func (m *SessionReportModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM session_reports WHERE id = $1`

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
