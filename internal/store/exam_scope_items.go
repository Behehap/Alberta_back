// internal/store/exam_scope_items.go
package store

import (
	"context"
	"database/sql"
	"time"
)

// ExamScopeItem defines a specific lesson that is part of an exam's scope.
type ExamScopeItem struct {
	ID            int64  `json:"id"`
	ExamID        int64  `json:"exam_id"`
	LessonID      int64  `json:"lesson_id"`
	TitleOverride string `json:"title_override,omitempty"`
}

// ExamScopeItemModel holds the database connection.
type ExamScopeItemModel struct {
	DB *sql.DB
}

// Insert adds a new lesson to an exam's scope.
func (m *ExamScopeItemModel) Insert(ctx context.Context, esi *ExamScopeItem) error {
	query := `
        INSERT INTO exam_scope_items (exam_id, lesson_id, title_override)
        VALUES ($1, $2, $3)
        RETURNING id`

	args := []any{esi.ExamID, esi.LessonID, esi.TitleOverride}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&esi.ID)
}

// GetAllForExam retrieves all scope items for a specific exam.
func (m *ExamScopeItemModel) GetAllForExam(ctx context.Context, examID int64) ([]*ExamScopeItem, error) {
	if examID < 1 {
		return nil, ErrorNotFound
	}

	query := `
        SELECT id, exam_id, lesson_id, title_override
        FROM exam_scope_items
        WHERE exam_id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, examID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*ExamScopeItem
	for rows.Next() {
		var esi ExamScopeItem
		err := rows.Scan(
			&esi.ID,
			&esi.ExamID,
			&esi.LessonID,
			&esi.TitleOverride,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, &esi)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// Delete removes a lesson from an exam's scope.
func (m *ExamScopeItemModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}

	query := `DELETE FROM exam_scope_items WHERE id = $1`

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
