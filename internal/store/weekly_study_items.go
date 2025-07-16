package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type WeeklyStudyItem struct {
	ID             int64     `json:"id"`
	WeeklyPlanID   int64     `json:"weekly_plan_id"`
	LessonID       int64     `json:"lesson_id"`
	IsCompleted    bool      `json:"is_completed"`
	CompletionDate time.Time `json:"completion_date,omitempty"`
}

type WeeklyStudyItemModel struct {
	DB *sql.DB
}

func (m *WeeklyStudyItemModel) Insert(ctx context.Context, wsi *WeeklyStudyItem) error {
	query := `
        INSERT INTO weekly_study_items (weekly_plan_id, lesson_id)
        VALUES ($1, $2)
        RETURNING id, is_completed`

	args := []any{wsi.WeeklyPlanID, wsi.LessonID}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&wsi.ID, &wsi.IsCompleted)
}

func (m *WeeklyStudyItemModel) GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*WeeklyStudyItem, error) {
	query := `
        SELECT id, weekly_plan_id, lesson_id, is_completed, completion_date
        FROM weekly_study_items
        WHERE weekly_plan_id = $1
        ORDER BY id`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, weeklyPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*WeeklyStudyItem
	for rows.Next() {
		var item WeeklyStudyItem
		var completionDate sql.NullTime
		err := rows.Scan(
			&item.ID,
			&item.WeeklyPlanID,
			&item.LessonID,
			&item.IsCompleted,
			&completionDate,
		)
		if err != nil {
			return nil, err
		}
		if completionDate.Valid {
			item.CompletionDate = completionDate.Time
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (m *WeeklyStudyItemModel) Update(ctx context.Context, wsi *WeeklyStudyItem) error {
	query := `
        UPDATE weekly_study_items
        SET is_completed = $1, completion_date = $2
        WHERE id = $3`

	var completionDate sql.NullTime
	if wsi.IsCompleted {
		completionDate.Time = time.Now()
		completionDate.Valid = true
	}

	args := []any{wsi.IsCompleted, completionDate, wsi.ID}
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

func (m *WeeklyStudyItemModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}
	query := `DELETE FROM weekly_study_items WHERE id = $1`

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

func (m *WeeklyStudyItemModel) Get(ctx context.Context, id int64) (*WeeklyStudyItem, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}
	query := `
        SELECT id, weekly_plan_id, lesson_id, is_completed, completion_date
        FROM weekly_study_items
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var item WeeklyStudyItem
	var completionDate sql.NullTime
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&item.ID,
		&item.WeeklyPlanID,
		&item.LessonID,
		&item.IsCompleted,
		&completionDate,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	if completionDate.Valid {
		item.CompletionDate = completionDate.Time
	}

	return &item, nil
}
