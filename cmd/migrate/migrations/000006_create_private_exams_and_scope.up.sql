-- 000006_create_private_exams_and_scope.up.sql

CREATE TABLE private_exams (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    exam_date DATE NOT NULL,
    organisation VARCHAR(255),
    target_grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id)
);

CREATE TABLE exam_scope_items (
    id SERIAL PRIMARY KEY,
    exam_id INT NOT NULL REFERENCES private_exams(id) ON DELETE CASCADE, -- References new 'private_exams' table
    lesson_id INT NOT NULL REFERENCES lessons(id),
    title_override VARCHAR(255)
);