-- scripts/seed.sql
-- This file is for inserting initial data into tables.
-- Table creation and schema changes are handled by migration files.

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

-- Step 6: Template rules (with new columns, assuming they are added by migrations)
INSERT INTO template_rules (template_id, book_id, default_frequency, scheduling_hints, consecutive_sessions, time_preference, priority_slot) VALUES
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 3, 'priority_first_block', FALSE, NULL, 'first'), -- Assuming 'priority_first_block' translates to priority_slot 'first'
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'عربی (دهم)'), 3, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 2, 'contiguous_pair', TRUE, NULL, NULL), -- Assuming 'contiguous_pair' translates to consecutive_sessions TRUE
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'تاریخ (دهم)'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'منطق (دهم)'), 2, NULL, FALSE, NULL, NULL);

-- Step 7: Seed book_roles to link books to curriculum
-- For "General" vs. "Special" books:
-- If major_id is NULL, it means the book applies to ALL majors for that grade.
-- If major_id is specified, it applies only to that specific major.
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'ادبیات فارسی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'عربی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'دین و زندگی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'جامعه شناسی (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'تاریخ (دهم)'), 'Core' ),
( (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'منطق (دهم)'), 'Core' );
