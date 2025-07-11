-- 000003_create_scheduling_tables.up.sql

CREATE TABLE IF NOT EXISTS unavailable_times (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    title VARCHAR(255),
    day_of_week INT NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    is_recurring BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS weekly_plans (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    start_date_of_week DATE NOT NULL,
    day_start_time TIME,
    max_study_time_hours_per_week INT,
    UNIQUE(student_id, start_date_of_week)
);

CREATE TABLE IF NOT EXISTS subject_frequencies (
    id SERIAL PRIMARY KEY,
    weekly_plan_id INT NOT NULL REFERENCES weekly_plans(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    frequency_per_week INT NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS daily_plans (
    id SERIAL PRIMARY KEY,
    weekly_plan_id INT NOT NULL REFERENCES weekly_plans(id) ON DELETE CASCADE,
    plan_date DATE NOT NULL,
    UNIQUE(weekly_plan_id, plan_date)
);