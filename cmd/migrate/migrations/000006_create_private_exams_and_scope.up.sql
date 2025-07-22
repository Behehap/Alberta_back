CREATE TABLE exam_schedules (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    exam_date DATE NOT NULL,
    organisation VARCHAR(255),
    target_grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id)
);

CREATE TABLE exam_scope_items (
    id SERIAL PRIMARY KEY,
    exam_id INT NOT NULL REFERENCES exam_schedules(id) ON DELETE CASCADE,
    lesson_id INT NOT NULL REFERENCES lessons(id),
    title_override VARCHAR(255)
);