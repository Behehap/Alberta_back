

CREATE TABLE IF NOT EXISTS weekly_study_items (
    id SERIAL PRIMARY KEY,
    weekly_plan_id INT NOT NULL REFERENCES weekly_plans(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id),
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    completion_date DATE
);

CREATE TABLE IF NOT EXISTS session_reports (
    id SERIAL PRIMARY KEY,

    weekly_study_item_id INT UNIQUE NOT NULL REFERENCES weekly_study_items(id) ON DELETE CASCADE,
    is_review BOOLEAN NOT NULL DEFAULT FALSE,
    num_tests INT,
    num_wrong_tests INT,
    session_score DECIMAL(5, 2),
    notes TEXT
);
