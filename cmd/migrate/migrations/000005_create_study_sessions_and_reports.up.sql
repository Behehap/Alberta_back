-- 000005_create_study_sessions_and_reports.up.sql

CREATE TABLE study_sessions (
    id SERIAL PRIMARY KEY,
    daily_plan_id INT NOT NULL REFERENCES daily_plans(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    is_completed BOOLEAN NOT NULL DEFAULT FALSE,
    completion_date DATE,
    start_time TIME NOT NULL, 
    end_time TIME NOT NULL    
);

CREATE TABLE session_reports (
    id SERIAL PRIMARY KEY,
    study_session_id INT UNIQUE NOT NULL REFERENCES study_sessions(id) ON DELETE CASCADE,
    is_review BOOLEAN NOT NULL DEFAULT FALSE,
    num_tests INT,
    num_wrong_tests INT,
    session_score DECIMAL(5, 2),
    notes TEXT
);