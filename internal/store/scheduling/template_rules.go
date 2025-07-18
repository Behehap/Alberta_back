package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type TemplateRule struct {
	ID                  int64          `json:"id"`
	TemplateID          int64          `json:"template_id"`
	BookID              int64          `json:"book_id"`
	DefaultFrequency    int            `json:"default_frequency"`
	SchedulingHints     sql.NullString `json:"scheduling_hints,omitempty"`
	ConsecutiveSessions sql.NullBool   `json:"consecutive_sessions,omitempty"`
	TimePreference      sql.NullString `json:"time_preference,omitempty"`
	PrioritySlot        sql.NullString `json:"priority_slot,omitempty"`
}

type TemplateRuleModel struct {
	DB *sql.DB
}

func (m *TemplateRuleModel) Insert(ctx context.Context, tr *TemplateRule) error {
	query := `
        INSERT INTO template_rules (template_id, book_id, default_frequency, scheduling_hints, consecutive_sessions, time_preference, priority_slot)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	args := []any{
		tr.TemplateID,
		tr.BookID,
		tr.DefaultFrequency,
		tr.SchedulingHints,
		tr.ConsecutiveSessions,
		tr.TimePreference,
		tr.PrioritySlot,
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&tr.ID)
}

func (m *TemplateRuleModel) Get(ctx context.Context, id int64) (*TemplateRule, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, template_id, book_id, default_frequency, scheduling_hints,
               consecutive_sessions, time_preference, priority_slot
        FROM template_rules
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var tr TemplateRule
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&tr.ID,
		&tr.TemplateID,
		&tr.BookID,
		&tr.DefaultFrequency,
		&tr.SchedulingHints,
		&tr.ConsecutiveSessions,
		&tr.TimePreference,
		&tr.PrioritySlot,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &tr, nil
}

func (m *TemplateRuleModel) GetAllForTemplate(ctx context.Context, templateID int64) ([]*TemplateRule, error) {
	if templateID < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, template_id, book_id, default_frequency, scheduling_hints,
               consecutive_sessions, time_preference, priority_slot
        FROM template_rules
        WHERE template_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, templateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []*TemplateRule
	for rows.Next() {
		var rule TemplateRule
		err := rows.Scan(
			&rule.ID,
			&rule.TemplateID,
			&rule.BookID,
			&rule.DefaultFrequency,
			&rule.SchedulingHints,
			&rule.ConsecutiveSessions,
			&rule.TimePreference,
			&rule.PrioritySlot,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, &rule)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}

func (m *TemplateRuleModel) Update(ctx context.Context, tr *TemplateRule) error {
	query := `
        UPDATE template_rules
        SET book_id = $1, default_frequency = $2, scheduling_hints = $3,
            consecutive_sessions = $4, time_preference = $5, priority_slot = $6
        WHERE id = $7 AND template_id = $8`

	args := []any{
		tr.BookID,
		tr.DefaultFrequency,
		tr.SchedulingHints,
		tr.ConsecutiveSessions,
		tr.TimePreference,
		tr.PrioritySlot,
		tr.ID,
		tr.TemplateID,
	}

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

func (m *TemplateRuleModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM template_rules WHERE id = $1`

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
