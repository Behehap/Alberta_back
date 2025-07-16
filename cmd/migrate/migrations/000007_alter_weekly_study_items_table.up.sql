-- 000007_alter_weekly_study_items_table.up.sql

-- Drop the foreign key constraint on lesson_id first
ALTER TABLE weekly_study_items DROP CONSTRAINT IF EXISTS weekly_study_items_lesson_id_fkey;

-- Drop the lesson_id column
ALTER TABLE weekly_study_items DROP COLUMN IF EXISTS lesson_id;

-- Add the new book_id column
ALTER TABLE weekly_study_items ADD COLUMN book_id INT NOT NULL;

-- Add a foreign key constraint for book_id
ALTER TABLE weekly_study_items ADD CONSTRAINT fk_book_id
FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE;