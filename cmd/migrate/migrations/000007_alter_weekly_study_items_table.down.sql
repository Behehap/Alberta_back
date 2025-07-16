-- 000007_alter_weekly_study_items_table.down.sql

-- Drop the foreign key constraint on book_id
ALTER TABLE weekly_study_items DROP CONSTRAINT IF EXISTS fk_book_id;

-- Drop the book_id column
ALTER TABLE weekly_study_items DROP COLUMN IF EXISTS book_id;

-- Re-add the lesson_id column
-- Note: You might need to adjust default values or allow NULLs if there was no NOT NULL constraint originally
-- and you had existing data that didn't have lessons. For simplicity, re-adding as NOT NULL here.
ALTER TABLE weekly_study_items ADD COLUMN lesson_id INT NOT NULL;

-- Re-add the foreign key constraint on lesson_id
ALTER TABLE weekly_study_items ADD CONSTRAINT weekly_study_items_lesson_id_fkey
FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE;