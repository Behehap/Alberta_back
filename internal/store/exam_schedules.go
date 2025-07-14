package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type ExamSchedule struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	ExamDate      time.Time `json:"exam_date"`
	Organisation  string    `json:"organisation,omitempty"`
	TargetGradeID int64     `json:"target_grade_id"`
	MajorID       int64     `json:"major_id"`
}

type ExamScheduleModel struct {
	DB *sql.DB
}

func (m *ExamScheduleModel) Insert(ctx context.Context, es *ExamSchedule) error {
	query := `
        INSERT INTO exam_schedules (title, exam_date, organisation, target_grade_id, major_id)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`

	args := []any{es.Title, es.ExamDate, es.Organisation, es.TargetGradeID, es.MajorID}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&es.ID)
}

func (m *ExamScheduleModel) Get(ctx context.Context, id int64) (*ExamSchedule, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, title, exam_date, organisation, target_grade_id, major_id
        FROM exam_schedules
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var es ExamSchedule
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&es.ID,
		&es.Title,
		&es.ExamDate,
		&es.Organisation,
		&es.TargetGradeID,
		&es.MajorID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &es, nil
}

func (m *ExamScheduleModel) GetAllForStudentCurriculum(ctx context.Context, gradeID, majorID int64) ([]*ExamSchedule, error) {
	query := `
        SELECT id, title, exam_date, organisation, target_grade_id, major_id
        FROM exam_schedules
        WHERE target_grade_id = $1 AND major_id = $2
        ORDER BY exam_date ASC`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, gradeID, majorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exams []*ExamSchedule
	for rows.Next() {
		var es ExamSchedule
		err := rows.Scan(
			&es.ID,
			&es.Title,
			&es.ExamDate,
			&es.Organisation,
			&es.TargetGradeID,
			&es.MajorID,
		)
		if err != nil {
			return nil, err
		}
		exams = append(exams, &es)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return exams, nil
}

func (m *ExamScheduleModel) Update(ctx context.Context, es *ExamSchedule) error {
	query := `
        UPDATE exam_schedules
        SET title = $1, exam_date = $2, organisation = $3, target_grade_id = $4, major_id = $5
        WHERE id = $6`

	args := []any{es.Title, es.ExamDate, es.Organisation, es.TargetGradeID, es.MajorID, es.ID}

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

func (m *ExamScheduleModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM exam_schedules WHERE id = $1`

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
