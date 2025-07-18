package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type ScheduleTemplate struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name"`
	TargetGradeID           int64  `json:"target_grade_id"`
	TargetMajorID           int64  `json:"target_major_id"`
	TotalStudyBlocksPerWeek int    `json:"total_study_blocks_per_week"`
}

type ScheduleTemplateModel struct {
	DB *sql.DB
}

func (m *ScheduleTemplateModel) Get(ctx context.Context, id int64) (*ScheduleTemplate, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, name, target_grade_id, target_major_id, total_study_blocks_per_week
        FROM schedule_templates
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var tpl ScheduleTemplate
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&tpl.ID,
		&tpl.Name,
		&tpl.TargetGradeID,
		&tpl.TargetMajorID,
		&tpl.TotalStudyBlocksPerWeek,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &tpl, nil
}

func (m *ScheduleTemplateModel) GetAll(ctx context.Context, gradeID, majorID int64) ([]*ScheduleTemplate, error) {
	query := `
        SELECT id, name, target_grade_id, target_major_id, total_study_blocks_per_week
        FROM schedule_templates
        WHERE target_grade_id = $1 AND target_major_id = $2
        ORDER BY name`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, gradeID, majorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var templates []*ScheduleTemplate
	for rows.Next() {
		var tpl ScheduleTemplate
		err := rows.Scan(
			&tpl.ID,
			&tpl.Name,
			&tpl.TargetGradeID,
			&tpl.TargetMajorID,
			&tpl.TotalStudyBlocksPerWeek,
		)
		if err != nil {
			return nil, err
		}
		templates = append(templates, &tpl)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return templates, nil
}
