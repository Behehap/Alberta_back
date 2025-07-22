// internal/store/book_roles.go
package store

import (
	"context"
	"database/sql"
	"time"
)

type BookRole struct {
	ID                   int64  `json:"id"`
	TargetStudentGradeID int64  `json:"target_student_grade_id"`
	MajorID              int64  `json:"major_id"`
	BookID               int64  `json:"book_id"`
	Role                 string `json:"role"`
}

type BookRoleModel struct {
	DB *sql.DB
}

func (m *BookRoleModel) GetAllForCurriculum(ctx context.Context, gradeID, majorID int64) ([]*BookRole, error) {
	query := `
        SELECT id, target_student_grade_id, major_id, book_id, role
        FROM book_roles
        WHERE target_student_grade_id = $1 AND major_id = $2`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, gradeID, majorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookRoles []*BookRole

	for rows.Next() {
		var bookRole BookRole
		err := rows.Scan(
			&bookRole.ID,
			&bookRole.TargetStudentGradeID,
			&bookRole.MajorID,
			&bookRole.BookID,
			&bookRole.Role,
		)
		if err != nil {
			return nil, err
		}
		bookRoles = append(bookRoles, &bookRole)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookRoles, nil
}
