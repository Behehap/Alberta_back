// internal/store/daily_plans.go
package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

// DailyPlan represents a specific day within a student's weekly plan.
type DailyPlan struct {
	ID           int64     `json:"id"`
	WeeklyPlanID int64     `json:"weekly_plan_id"`
	PlanDate     time.Time `json:"plan_date"`
}

// DailyPlanModel holds the database connection.
type DailyPlanModel struct {
	DB *sql.DB
}

// Insert creates a new daily plan record.
func (m *DailyPlanModel) Insert(ctx context.Context, dp *DailyPlan) error {
	query := `
        INSERT INTO daily_plans (weekly_plan_id, plan_date)
        VALUES ($1, $2)
        RETURNING id`

	args := []any{dp.WeeklyPlanID, dp.PlanDate}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&dp.ID)
}

// GetByDate retrieves a daily plan for a specific weekly plan and date.
func (m *DailyPlanModel) GetByDate(ctx context.Context, weeklyPlanID int64, planDate time.Time) (*DailyPlan, error) {
	query := `
        SELECT id, weekly_plan_id, plan_date
        FROM daily_plans
        WHERE weekly_plan_id = $1 AND plan_date = $2`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var dp DailyPlan
	err := m.DB.QueryRowContext(ctx, query, weeklyPlanID, planDate).Scan(
		&dp.ID,
		&dp.WeeklyPlanID,
		&dp.PlanDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &dp, nil
}

// GetAllForWeeklyPlan retrieves all daily plans associated with a weekly plan.
func (m *DailyPlanModel) GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*DailyPlan, error) {
	query := `
        SELECT id, weekly_plan_id, plan_date
        FROM daily_plans
        WHERE weekly_plan_id = $1
        ORDER BY plan_date ASC`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, weeklyPlanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plans []*DailyPlan
	for rows.Next() {
		var dp DailyPlan
		err := rows.Scan(
			&dp.ID,
			&dp.WeeklyPlanID,
			&dp.PlanDate,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, &dp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return plans, nil
}
