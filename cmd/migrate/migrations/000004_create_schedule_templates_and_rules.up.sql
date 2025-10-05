-- 000004_create_schedule_templates_and_rules.up.sql

CREATE TABLE schedule_templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    target_grade_id INT NOT NULL REFERENCES grades(id),
    target_major_id INT NOT NULL REFERENCES majors(id),
    total_study_blocks_per_week INT NOT NULL
);

CREATE TABLE template_rules (
    id SERIAL PRIMARY KEY,
    template_id INT NOT NULL REFERENCES schedule_templates(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    default_frequency INT NOT NULL,
    scheduling_hints TEXT, -- Keeping original string for flexibility
    consecutive_sessions BOOLEAN NOT NULL DEFAULT FALSE,
    time_preference VARCHAR(10), -- 'morning' or 'afternoon'
    priority_slot VARCHAR(20) -- 'first', 'any'
);


-- Add after existing tables
CREATE TABLE template_subject_weights (
    id SERIAL PRIMARY KEY,
    template_id INT NOT NULL REFERENCES schedule_templates(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    weight DECIMAL(3,2) NOT NULL DEFAULT 1.0,
    UNIQUE(template_id, book_id)
);