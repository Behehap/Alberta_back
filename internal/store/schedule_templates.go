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
