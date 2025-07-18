package store

import (
	"context"
	"database/sql"
	"time"
)

type TemplateRule struct {
	ID                  int64          `json:"id"`
	TemplateID          int64          `json:"template_id"`
	BookID              int64          `json:"book_id"`
	DefaultFrequency    int            `json:"default_frequency"`
	SchedulingHints     string         `json:"scheduling_hints,omitempty"`
	ConsecutiveSessions bool           `json:"consecutive_sessions"`
	TimePreference      sql.NullString `json:"time_preference,omitempty"`
	PrioritySlot        sql.NullString `json:"priority_slot,omitempty"`
}

type TemplateRuleModel struct {
	DB *sql.DB
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
