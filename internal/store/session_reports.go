package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type SessionReport struct {
	ID int64 `json:"id"`

	StudySessionID int64   `json:"study_session_id"`
	IsReview       bool    `json:"is_review"`
	NumTests       int     `json:"num_tests,omitempty"`
	NumWrongTests  int     `json:"num_wrong_tests,omitempty"`
	SessionScore   float64 `json:"session_score,omitempty"`
	Notes          string  `json:"notes,omitempty"`
}

type SessionReportModel struct {
	DB *sql.DB
}

func (m *SessionReportModel) Insert(ctx context.Context, sr *SessionReport) error {
	query := `
        INSERT INTO session_reports (study_session_id, is_review, num_tests, num_wrong_tests, session_score, notes)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	args := []any{sr.StudySessionID, sr.IsReview, sr.NumTests, sr.NumWrongTests, sr.SessionScore, sr.Notes}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&sr.ID)
}

func (m *SessionReportModel) GetForStudySession(ctx context.Context, studySessionID int64) (*SessionReport, error) {

	if studySessionID < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, study_session_id, is_review, num_tests, num_wrong_tests, session_score, notes
        FROM session_reports
        WHERE study_session_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var sr SessionReport

	err := m.DB.QueryRowContext(ctx, query, studySessionID).Scan(
		&sr.ID,
		&sr.StudySessionID,
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

func (m *SessionReportModel) Update(ctx context.Context, sr *SessionReport) error {
	query := `
        UPDATE session_reports
        SET is_review = $1, num_tests = $2, num_wrong_tests = $3, session_score = $4, notes = $5
        WHERE id = $6`

	args := []any{sr.IsReview, sr.NumTests, sr.NumWrongTests, sr.SessionScore, sr.Notes, sr.ID}

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
