-- 000002_create_students_table.up.sql

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(50),
    grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id)
);