

CREATE TABLE IF NOT EXISTS schedule_templates (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    target_grade_id INT NOT NULL REFERENCES grades(id),
    target_major_id INT NOT NULL REFERENCES majors(id),
    total_study_blocks_per_week INT NOT NULL
);

CREATE TABLE IF NOT EXISTS template_rules (
    id SERIAL PRIMARY KEY,
    template_id INT NOT NULL REFERENCES schedule_templates(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    default_frequency INT NOT NULL,
    scheduling_hints TEXT
);
