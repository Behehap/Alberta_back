-- 000005_create_exam_tables.up.sql

CREATE TABLE IF NOT EXISTS exams (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    exam_date DATE NOT NULL,
    organisation VARCHAR(255),
    target_grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id)
);

CREATE TABLE IF NOT EXISTS exam_scope_items (
    id SERIAL PRIMARY KEY,
    exam_id INT NOT NULL REFERENCES exams(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id),
    title_override VARCHAR(255)
);