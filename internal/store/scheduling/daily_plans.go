package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type DailyPlan struct {
	ID           int64     `json:"id"`
	WeeklyPlanID int64     `json:"weekly_plan_id"`
	PlanDate     time.Time `json:"plan_date"`
}

type DailyPlanModel struct {
	DB *sql.DB
}

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

func (m *DailyPlanModel) Get(ctx context.Context, id int64) (*DailyPlan, error) {
	if id < 1 {
		return nil, ErrorNotFound
	}
	query := `
        SELECT id, weekly_plan_id, plan_date
        FROM daily_plans
        WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var dp DailyPlan
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
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

func (m *DailyPlanModel) GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*DailyPlan, error) {
	query := `
        SELECT id, weekly_plan_id, plan_date
        FROM daily_plans
        WHERE weekly_plan_id = $1
        ORDER BY plan_date`

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

func (m *DailyPlanModel) Delete(ctx context.Context, id int64) error {
	if id < 1 {
		return ErrorNotFound
	}
	query := `DELETE FROM daily_plans WHERE id = $1`

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

func (m *DailyPlanModel) GetByWeeklyPlanAndDate(ctx context.Context, weeklyPlanID int64, planDate time.Time) (*DailyPlan, error) {
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
