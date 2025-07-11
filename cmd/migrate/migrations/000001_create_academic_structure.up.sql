-- 000001_create_academic_structure.up.sql

CREATE TABLE IF NOT EXISTS grades (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS majors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    inherent_grade_level_id INT NOT NULL REFERENCES grades(id)
);

CREATE TABLE IF NOT EXISTS lessons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    estimated_study_time_minutes INT
);

CREATE TABLE IF NOT EXISTS book_roles (
    id SERIAL PRIMARY KEY,
    target_student_grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id),
    book_id INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    role VARCHAR(255) NOT NULL
);