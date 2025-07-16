-- Step 0: Create all tables
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

CREATE TABLE IF NOT EXISTS students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(50),
    grade_id INT NOT NULL REFERENCES grades(id),
    major_id INT NOT NULL REFERENCES majors(id)
);

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

-- Step 1: Clean out all tables and reset IDs
TRUNCATE TABLE
    grades,
    majors,
    books,
    lessons,
    book_roles,
    students,
    unavailable_times,
    weekly_plans,
    subject_frequencies,
    weekly_study_items,
    session_reports,
    exams,
    exam_scope_items,
    schedule_templates,
    template_rules
RESTART IDENTITY CASCADE;

-- Step 2: Seed grades
INSERT INTO grades (name) VALUES
('دهم'),
('یازدهم'),
('دوازدهم');

-- Step 3: Seed majors
INSERT INTO majors (name) VALUES
('علوم تجربی'),
('ریاضی فیزیک'),
('علوم انسانی');

-- Step 4: Seed books and lessons for 10th Humanities
-- Persian Literature
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ادبیات فارسی (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 'درس اول: چشمه و سنگ'),
((SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 'درس دوم: از آموختن، ننگ مدار');

-- Arabic
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی (دهم)'), 'الدرس الأول'),
((SELECT id FROM books WHERE title = 'عربی (دهم)'), 'الدرس الثانی');

-- Religion
INSERT INTO books (title, inherent_grade_level_id) VALUES ('دین و زندگی (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 'درس اول: هدف زندگی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 'درس دوم: پر پرواز');

-- Sociology
INSERT INTO books (title, inherent_grade_level_id) VALUES ('جامعه شناسی (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 'درس اول: کنش های ما'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 'درس دوم: پدیده های اجتماعی');

-- History
INSERT INTO books (title, inherent_grade_level_id) VALUES ('تاریخ (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'تاریخ (دهم)'), 'درس اول: تاریخ و تاریخ نگاری');

-- Logic
INSERT INTO books (title, inherent_grade_level_id) VALUES ('منطق (دهم)', 1);
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'منطق (دهم)'), 'درس اول: منطق، ترازوی اندیشه');

-- Step 5: Insert schedule template
INSERT INTO schedule_templates (name, target_grade_id, target_major_id, total_study_blocks_per_week) VALUES
('دهم انسانی - ۲۴ بلوک', 1, 3, 24);

-- Step 6: Template rules
INSERT INTO template_rules (template_id, book_id, default_frequency, scheduling_hints) VALUES
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 3, 'priority_first_block'),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'عربی (دهم)'), 3, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 2, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 2, 'contiguous_pair'),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'تاریخ (دهم)'), 2, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'منطق (دهم)'), 2, NULL);


-- Step 7: Seed book_roles to link books to curriculum
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'عربی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'تاریخ (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'منطق (دهم)'), 'Core' );