-- scripts/seed.sql
-- This file is for inserting initial data into tables.
-- Table creation and schema changes are handled by migration files.

-- Step 1: Clean existing data (Optional but recommended for a clean seed)
DELETE FROM session_reports;
DELETE FROM study_sessions;
DELETE FROM daily_plans;
DELETE FROM subject_frequencies;
DELETE FROM weekly_plans;
DELETE FROM unavailable_times;
DELETE FROM exam_scope_items;
DELETE FROM exam_schedules;
DELETE FROM template_rules;
DELETE FROM schedule_templates;
DELETE FROM book_roles;
DELETE FROM lessons;
DELETE FROM books;
DELETE FROM students; -- Delete students before grades/majors
DELETE FROM majors;
DELETE FROM grades;


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


-- =================================================================
--                              پایه دهم
-- =================================================================

-- Step 4.1: Seed Books & Lessons for 10th Grade (General - عمومی)

-- فارسی (۱)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فارسی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس اول: چشمه'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس دوم: از آموختن، ننگ مدار'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس سوم: پاسداری از حقیقت'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس چهارم: بیداد ظالمان'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس پنجم: مهر و وفا'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس ششم: جمال و کمال'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس هفتم: سفر به بصره'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس هشتم: کلاس نقّاشی'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس نهم: دریادلان صف شکن'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس دهم: خاک آزادگان'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس یازدهم: رستم و اشکبوس'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس دوازدهم: گُردآفرید'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس سیزدهم: طوطی و بقال'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس چهاردهم: سپیده دم'),
((SELECT id FROM books WHERE title = 'فارسی (۱)'), 'درس پانزدهم: عظمتِ نگاه');

-- عربی، زبان قرآن (۱) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان قرآن (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الْأَوَّلُ: ذاكَ هوَ اللّٰهُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الثّاني: اَلْمَواعِظُ الْعَدَديَّةُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الثّالِثُ: مَطَرُ السَّمَكِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الرّابِعُ: اَلتَّعايُشُ السِّلْميُّ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الْخامِسُ: هذا خَلْقُ اللّٰهِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ السّادِسُ: ...فَاعْلَمْ أَنَّهُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ السّابِعُ: ...صِناعَةُ التَّلْميعِ في الْأَدَبِ الْفارِسيِّ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'اَلدَّرْسُ الثّامِنُ: ...صُورَةٌ مِنَ الطَّبيعَةِ');

-- دین و زندگی (۱)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('دین و زندگی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس اول: هدف زندگی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس دوم: پر پرواز'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس سوم: پنجره‌ای به روشنایی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس چهارم: آینده روشن'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس پنجم: منزلگاه بعد'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس ششم: واقعه بزرگ'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس هفتم: فرجام کار'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس هشتم: آهنگ سفر'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس نهم: دوستی با خدا'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس دهم: یاری از نماز و روزه'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس یازدهم: فضیلت آراستگی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'درس دوازدهم: زیبایی پوشیدگی');

-- انگلیسی (۱)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('انگلیسی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Lesson 1: Saving Nature'),
((SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Lesson 2: Wonders of Creation'),
((SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Lesson 3: The Value of Knowledge'),
((SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Lesson 4: A Great Inventor');

-- نگارش (۱)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('نگارش (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس اول: پرورش موضوع'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس دوم: عینک نوشتن'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس سوم: نوشته های عینی'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس چهارم: نوشته های ذهنی (1)'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس پنجم: نوشته های ذهنی (2)'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس ششم: سنجش و مقایسه'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس هفتم: ناسازی معنایی یا تضاد مفاهیم'),
((SELECT id FROM books WHERE title = 'نگارش (۱)'), 'درس هشتم: نوشته های داستان گونه');

-- آمادگی دفاعی
INSERT INTO books (title, inherent_grade_level_id) VALUES ('آمادگی دفاعی', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس اول: امنیت و تهدید'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس دوم: اقتدار دفاعی'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس سوم: انقلاب اسلامی'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس چهارم: آشنایی با بسیج'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس پنجم: علوم و معارف دفاع مقدس'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس ششم: الگوها و اسوه های پایداری'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس هفتم: آشنایی با نیروهای مسلح'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس هشتم: جنگ نرم'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس نهم: نظام جمع و شیوه های رزم انفرادی'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس دهم: آمادگی در برابر زلزله'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس یازدهم: پدافند غیرعامل'),
((SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'درس دوازدهم: ایمنی و پیشگیری');


-- Step 4.2: Seed Books & Lessons for 10th Grade (Specialized - تخصصی)

-- شیمی (۱) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('شیمی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'شیمی (۱)'), 'فصل اول: کیهان زادگاه الفبای هستی'),
((SELECT id FROM books WHERE title = 'شیمی (۱)'), 'فصل دوم: ردِّپای گازها در زندگی'),
((SELECT id FROM books WHERE title = 'شیمی (۱)'), 'فصل سوم: آب، آهنگ زندگی');

-- ریاضی (۱) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ریاضی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل اول: مجموعه، الگو و دنباله'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل دوم: مثلثات'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل سوم: توان‌های گویا و عبارت‌های جبری'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل چهارم: معادله‌ها و نامعادله‌ها'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل پنجم: تابع'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل ششم: شمارش، بدون شمردن'),
((SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'فصل هفتم: آمار و احتمال');

-- فیزیک (۱) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۱) - ریاضی', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'فصل اول: فیزیک و اندازه‌گیری'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'فصل دوم: ویژگی‌های فیزیکی مواد'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'فصل سوم: کار، انرژی و توان'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'فصل چهارم: دما و گرما'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'فصل پنجم: ترمودینامیک');

-- هندسه (۱) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('هندسه (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'هندسه (۱)'), 'فصل اول: ترسیم‌های هندسی و استدلال'),
((SELECT id FROM books WHERE title = 'هندسه (۱)'), 'فصل دوم: قضیه تالس، تشابه و کاربردهای آن'),
((SELECT id FROM books WHERE title = 'هندسه (۱)'), 'فصل سوم: چندضلعی‌ها'),
((SELECT id FROM books WHERE title = 'هندسه (۱)'), 'فصل چهارم: تجسم فضایی');

-- فیزیک (۱) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۱) - تجربی', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۱) - تجربی'), 'فصل اول: فیزیک و اندازه‌گیری'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - تجربی'), 'فصل دوم: ویژگی‌های فیزیکی مواد'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - تجربی'), 'فصل سوم: کار، انرژی و توان'),
((SELECT id FROM books WHERE title = 'فیزیک (۱) - تجربی'), 'فصل چهارم: دما و گرما');

-- زیست شناسی (۱) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('زیست شناسی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل اول: دنیای زنده'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل دوم: گوارش و جذب مواد'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل سوم: تبادلات گازی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل چهارم: گردش مواد در بدن'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل پنجم: تنظیم اسمزی و دفع مواد زائد'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل ششم: از یاخته تا گیاه'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'فصل هفتم: جذب و انتقال مواد در گیاهان');

-- عربی، زبان تخصصی رشته انسانی (۱)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان تخصصی رشته انسانی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'اَلدَّرْسُ الْأَوَّلُ: أَنْواعُ الْفِعْلِ الثُّلاثیِّ الْمُجَرَّدِ وَ الْمَزیدِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'اَلدَّرْسُ الثّاني: اِسْمُ الْفاعِلِ وَ اسْمُ الْمَفْعولِ وَ اسْمُ الْمُبالَغَةِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'اَلدَّرْسُ الثّالِثُ: اَلْجُمَلُ الْاِسْمیَّةُ وَ الْفِعْلیَّةُ وَ أَنْواعُ الْإعْرابِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'اَلدَّرْسُ الرّابِعُ: إِعْرابُ الْفِعْلِ الْمُضارِعِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'اَلدَّرْسُ الْخامِسُ: اَلْفِعْلُ الْمَجْهولُ وَ الْمَعْلومُ');

-- علوم و فنون ادبی (۱) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('علوم و فنون ادبی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس اول: مبانی تحلیل متن'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس دوم: سازه‌ها و عوامل تأثیرگذار در شعر فارسی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس سوم: واج‌آرایی، واژه‌آرایی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس چهارم: تاریخ ادبیات پیش از اسلام و قرن‌های اولیۀ هجری'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس پنجم: هماهنگی پاره‌های کلام'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس ششم: سجع و انواع آن'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس هفتم: سبک و سبک‌شناسی (سبک خراسانی)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس هشتم: وزن شعر فارسی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس نهم: موازنه و ترصیع'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس دهم: زبان و ادبیات در سده‌های پنجم و ششم و ویژگی‌های سبک عراقی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس یازدهم: قافیه'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'درس دوازدهم: جناس و انواع آن');

-- منطق (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('منطق', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'منطق'), 'درس اول: منطق، ترازوی اندیشه'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس دوم: لفظ و معنا'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس سوم: مفهوم و مصداق'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس چهارم: اقسام تعریف و شرایط آن'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس پنجم: اقسام استدلال استقرایی'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس ششم: قضیۀ حملی'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس هفتم: احکام قضایا'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس هشتم: قیاس اقترانی'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس نهم: قضیۀ شرطی و قیاس استثنایی'),
((SELECT id FROM books WHERE title = 'منطق'), 'درس دهم: سنجشگری در تفکر');

-- اقتصاد (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('اقتصاد', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش اول: آشنایی با اقتصاد'),
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش دوم: تولید'),
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش سوم: بازار'),
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش چهارم: آشنایی با شاخص‌های اقتصادی'),
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش پنجم: پول و بانک'),
((SELECT id FROM books WHERE title = 'اقتصاد'), 'بخش ششم: اقتصاد ایران');

-- جامعه شناسی (۱) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('جامعه شناسی (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس اول: کنش‌های ما'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس دوم: پدیده‌های اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس سوم: جهان اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس چهارم: تشریح جهان اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس پنجم: معنای زندگی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس ششم: قدرت اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس هفتم: نابرابری اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس هشتم: ارزیابی جهان‌های اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس نهم: جهان متجدد'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس دهم: شناخت علمی جهان اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس یازدهم: تحولات هویتی جهان اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس دوازدهم: تحولات خانواده'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس سیزدهم: تحولات اقتصادی جهان اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'درس چهاردهم: تحولات سیاسی و اجتماعی جهان اجتماعی');

-- تاریخ (۱) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('تاریخ (۱)', (SELECT id FROM grades WHERE name = 'دهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس اول: تاریخ و تاریخ‌نگاری'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس دوم: تاریخ؛ زمان و مکان'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس سوم: باستان‌شناسی؛ در جستجوی ردپای انسان'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس چهارم: پیدایش تمدن؛ بین‌النهرین و مصر'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس پنجم: هند و چین'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس ششم: یونان و روم'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس هفتم: مطالعه و کاوش در گذشته‌های دور'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس هشتم: سپیده‌دمان تمدن ایرانی'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس نهم: از ورود آریایی‌ها تا پایان هخامنشیان'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس دهم: اشکانیان و ساسانیان'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس یازدهم: آیین کشورداری'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس دوازدهم: جامعه و خانواده'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس سیزدهم: اقتصاد و معیشت'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس چهاردهم: دین و اعتقادات'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس پانزدهم: فرهنگ و هنر'),
((SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'درس شانزدهم: میراث فرهنگی ایران باستان');


-- =================================================================
--                              پایه یازدهم
-- =================================================================

-- Step 5.1: Seed Books & Lessons for 11th Grade (General - عمومی)

-- فارسی (۲)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فارسی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس اول: نیکی'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس دوم: قاضی بُست'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس سوم: در کوی عاشقان'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس چهارم: درس آزاد (ادبیات بومی ۱)'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس پنجم: آغازگری تنها'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس ششم: پرورده عشق'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس هفتم: باران محبّت'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس هشتم: در امواج سند'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس نهم: آغاز و انجام'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس دهم: پرواز'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس یازدهم: کاوه دادخواه'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس دوازدهم: درس آزاد (ادبیات بومی ۲)'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس سیزدهم: خوان هشتم'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس چهاردهم: حمله حیدری'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس پانزدهم: کبوتر طوق‌دار'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس شانزدهم: قصّۀ عینکم'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس هفدهم: خاموشی دریا'),
((SELECT id FROM books WHERE title = 'فارسی (۲)'), 'درس هجدهم: شکوه چشمانت');

-- عربی، زبان قرآن (۲) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان قرآن (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ الْأَوَّلُ: مِنْ آیاتِ الْأَخْلاقِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ الثّاني: اَلصِّدْقُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ الثّالِثُ: عَجائِبُ الْمَخلوقاتِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ الرّابِعُ: آدابُ الْکَلامِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ الْخامِسُ: اَلصُّوَرُ الْجَمالیَّةُ فی الطَّبیعَةِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ السّادِسُ: اَلرَّجاءُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'اَلدَّرْسُ السّابِعُ: تَأثیرُ اللُّغَةِ الْفارِسیَّةِ عَلَی اللُّغَةِ الْعَرَبیَّةِ');

-- دین و زندگی (۲)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('دین و زندگی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس اول: هدایت الهی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس دوم: تداوم هدایت'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس سوم: معجزه جاویدان'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس چهارم: مسئولیت‌های پیامبر'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس پنجم: امامت، تداوم رسالت'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس ششم: پیشوایان اسوه'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس هفتم: وضعیت فرهنگی، اجتماعی و سیاسی عصر ائمه'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس هشتم: احیای ارزش‌های راستین'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس نهم: عصر غیبت'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس دهم: مرجعیت و ولایت فقیه'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس یازدهم: عزت نفس'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'درس دوازدهم: پیوند مقدس');

-- انگلیسی (۲)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('انگلیسی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Lesson 1: Understanding People'),
((SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Lesson 2: A Healthy Lifestyle'),
((SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Lesson 3: Art and Culture');

-- نگارش (۲)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('نگارش (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس اول: اجزای نوشته؛ ساختار و محتوا'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس دوم: گسترش محتوا (1)؛ زمان و مکان'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس سوم: گسترش محتوا (2)؛ شخصیت'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس چهارم: گسترش محتوا (3)؛ گفت و گو'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس پنجم: سفرنامه'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس ششم: گزارش نویسی'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس هفتم: کوتاه نویسی'),
((SELECT id FROM books WHERE title = 'نگارش (۲)'), 'درس هشتم: بازآفرینی و نقد');


-- Step 5.2: Seed Books & Lessons for 11th Grade (Specialized - تخصصی)

-- شیمی (۲) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('شیمی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'شیمی (۲)'), 'فصل اول: قدر هدایای زمینی را بدانیم'),
((SELECT id FROM books WHERE title = 'شیمی (۲)'), 'فصل دوم: در پی غذای سالم'),
((SELECT id FROM books WHERE title = 'شیمی (۲)'), 'فصل سوم: پوشاک، نیازی پایان ناپذیر');

-- حسابان (۱) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('حسابان (۱)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'حسابان (۱)'), 'فصل اول: جبر و معادله'),
((SELECT id FROM books WHERE title = 'حسابان (۱)'), 'فصل دوم: تابع'),
((SELECT id FROM books WHERE title = 'حسابان (۱)'), 'فصل سوم: توابع نمایی و لگاریتمی'),
((SELECT id FROM books WHERE title = 'حسابان (۱)'), 'فصل چهارم: مثلثات'),
((SELECT id FROM books WHERE title = 'حسابان (۱)'), 'فصل پنجم: حد و پیوستگی');

-- آمار و احتمال (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('آمار و احتمال', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'آمار و احتمال'), 'فصل اول: آشنایی با مبانی ریاضیات'),
((SELECT id FROM books WHERE title = 'آمار و احتمال'), 'فصل دوم: احتمال'),
((SELECT id FROM books WHERE title = 'آمار و احتمال'), 'فصل سوم: آمار توصیفی'),
((SELECT id FROM books WHERE title = 'آمار و احتمال'), 'فصل چهارم: آمار استنباطی');

-- هندسه (۲) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('هندسه (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'هندسه (۲)'), 'فصل اول: دایره'),
((SELECT id FROM books WHERE title = 'هندسه (۲)'), 'فصل دوم: تبدیل‌های هندسی و کاربردها'),
((SELECT id FROM books WHERE title = 'هندسه (۲)'), 'فصل سوم: روابط طولی در مثلث');

-- فیزیک (۲) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۲) - ریاضی', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۲) - ریاضی'), 'فصل اول: الکتریسیتۀ ساکن'),
((SELECT id FROM books WHERE title = 'فیزیک (۲) - ریاضی'), 'فصل دوم: جریان الکتریکی و مدارهای جریان مستقیم'),
((SELECT id FROM books WHERE title = 'فیزیک (۲) - ریاضی'), 'فصل سوم: مغناطیس'),
((SELECT id FROM books WHERE title = 'فیزیک (۲) - ریاضی'), 'فصل چهارم: القای الکترومغناطیسی و جریان متناوب');

-- ریاضی (۲) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ریاضی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل اول: هندسۀ تحلیلی و جبر'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل دوم: هندسه'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل سوم: تابع'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل چهارم: مثلثات'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل پنجم: توابع نمایی و لگاریتمی'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'ف فصل ششم: حد و پیوستگی'),
((SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'فصل هفتم: آمار و احتمال');

-- فیزیک (۲) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۲) - تجربی', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۲) - تجربی'), 'فصل اول: الکتریسیتۀ ساکن'),
((SELECT id FROM books WHERE title = 'فیزیک (۲) - تجربی'), 'فصل دوم: جریان الکتریکی و مدارهای جریان مستقیم'),
((SELECT id FROM books WHERE title = 'فیزیک (۲) - تجربی'), 'فصل سوم: مغناطیس و القای الکترومغناطیسی');

-- زیست شناسی (۲) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('زیست شناسی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل اول: تنظیم عصبی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل دوم: حواس'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل سوم: دستگاه حرکتی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل چهارم: تنظیم شیمیایی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل پنجم: ایمنی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل ششم: تقسیم یاخته'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل هفتم: تولید مثل'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل هشتم: تولید مثل نهاندانگان'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'فصل نهم: پاسخ گیاهان به محرک‌ها');

-- عربی، زبان تخصصی رشته انسانی (۲)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان تخصصی رشته انسانی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'اَلدَّرْسُ الْأَوَّلُ: اِسْمُ التَّفْضیلِ وَ اِسْمُ الْمَکانِ وَ اِسْمُ الزَّمانِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'اَلدَّرْسُ الثّاني: أُسْلوبُ الشَّرْطِ وَ أَدَواتُهُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'اَلدَّرْسُ الثّالِثُ: اَلْمَفاعیلُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'اَلدَّرْسُ الرّابِعُ: اَلْحالُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'اَلدَّرْسُ الْخامِسُ: اَلتَّمییزُ');

-- علوم و فنون ادبی (۲) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('علوم و فنون ادبی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس اول: تاریخ ادبیات در قرن‌های هفتم، هشتم و نهم'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس دوم: پایه‌های آوایی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس سوم: تشبیه'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس چهارم: سبک‌شناسی قرن‌های هفتم، هشتم و نهم (سبک عراقی)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس پنجم: پایه‌های آوایی همسان (۱)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس ششم: مجاز'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس هفتم: تاریخ ادبیات در قرن‌های دهم و یازدهم'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس هشتم: پایه‌های آوایی همسان (۲)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس نهم: استعاره'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس دهم: سبک‌شناسی قرن‌های دهم و یازدهم (سبک هندی)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس یازدهم: پایه‌های آوایی ناهمسان'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'درس دوازدهم: کنایه');

-- فلسفه (۱) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فلسفه (۱)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس اول: چیستی فلسفه'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس دوم: ریشه و شاخه‌های فلسفه'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس سوم: فلسفه و زندگی'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس چهارم: آغاز تاریخی فلسفه'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس پنجم: زندگی بر اساس اندیشه'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس ششم: امکان شناخت'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس هفتم: ابزارهای شناخت'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس هشتم: نگاهی به تاریخچۀ معرفت'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس نهم: چیستی انسان (۱)'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس دهم: چیستی انسان (۲)'),
((SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'درس یازدهم: انسان، موجود اخلاق‌گرا');

-- روان شناسی (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('روان شناسی', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس اول: روان‌شناسی؛ تعریف و روش مورد مطالعه'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس دوم: روان‌شناسی رشد'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس سوم: احساس، توجه، ادراک'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس چهارم: حافظه و علل فراموشی'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس پنجم: تفکر (۱) حل مسئله'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس ششم: تفکر (۲) تصمیم‌گیری'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس هفتم: انگیزش و هیجان'),
((SELECT id FROM books WHERE title = 'روان شناسی'), 'درس هشتم: روان‌شناسی سلامت');

-- جامعه شناسی (۲) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('جامعه شناسی (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس اول: فرهنگ جهانی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس دوم: فرهنگ جهانی (۲)'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس سوم: نمونه‌های فرهنگ جهانی (۱)'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس چهارم: نمونه‌های فرهنگ جهانی (۲)'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس پنجم: باورها و ارزش‌های بنیادین فرهنگ غرب'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس ششم: چگونگی تکوین فرهنگ معاصر غرب'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس هفتم: جامعۀ جهانی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس هشتم: تحولات نظام جهانی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس نهم: جهان دو قطبی و جهان تک‌قطبی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس دهم: جنگ‌ها و تقابل‌های جهانی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس یازدهم: بحران‌های اقتصادی و زیست محیطی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس دوازدهم: بحران‌های معرفتی و معنوی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس سیزدهم: سرنوشت جوامع'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس چهاردهم: افق آینده'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'درس پانزدهم: چشم‌اندازهای امیدبخش');

-- تاریخ (۲) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('تاریخ (۲)', (SELECT id FROM grades WHERE name = 'یازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس اول: منابع پژوهش در تاریخ اسلام و ایرانِ دوران اسلامی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس دوم: جهان در آستانۀ بعثت'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس سوم: ظهور اسلام و گسترش آن در شبه جزیرۀ عربستان'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس چهارم: خلافت؛ تحولات سیاسی، اجتماعی و اقتصادی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس پنجم: خلافت اموی و مروانی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس ششم: خلافت عباسی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس هفتم: جهان اسلام و میراث فرهنگی آن'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس هشتم: ایران در قرون نخستین اسلامی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس نهم: ظهور و گسترش تمدن ایرانی - اسلامی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس دهم: ایران در دوران غزنوی، سلجوقی و خوارزمشاهی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس یازدهم: حکومت، جامعه و اقتصاد در ایرانِ عصر سلجوقی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس دوازدهم: فرهنگ و هنر در عصر سلجوقی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس سیزدهم: یورش مغولان و پیامدهای آن'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس چهاردهم: ایران در عصر ایلخانان'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس پانزدهم: برآمدن تیمور و ظهور دولت صفوی'),
((SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'درس شانزدهم: تاریخ‌نگاری و گونه‌های آن');


-- =================================================================
--                              پایه دوازدهم
-- =================================================================

-- Step 6.1: Seed Books & Lessons for 12th Grade (General - عمومی)

-- فارسی (۳)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فارسی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس اول: شکرِ نعمت'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس دوم: مست و هشیار'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس سوم: آزادی'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس چهارم: درس آزاد (ادبیات بومی ۱)'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس پنجم: دماوندیه'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس ششم: نی‌نامه'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس هفتم: در حقیقتِ عشق'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس هشتم: از پاریز تا پاریس'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس نهم: کویر'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس دهم: فصل شکوفایی'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس یازدهم: آن شب عزیز'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس دوازدهم: گذران روزگار'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس سیزدهم: خوانِ آخر'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس چهاردهم: مرغِ گرفتار'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس پانزدهم: بچه‌های آسمان'),
((SELECT id FROM books WHERE title = 'فارسی (۳)'), 'درس شانزدهم: خندۀ تو');

-- عربی، زبان قرآن (۳) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان قرآن (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'اَلدَّرْسُ الْأَوَّلُ: اَلدّینُ وَ التَّدَیُّنُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'اَلدَّرْسُ الثّاني: مَکّةُ الْمُکَرَّمَةُ وَ الْمَدینَةُ الْمُنَوَّرَةُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'اَلدَّرْسُ الثّالِثُ: اَلْکُتُبُ طَعامُ الْفِکْرِ'),
((SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'اَلدَّرْسُ الرّابِعُ: اَلْفَرَزْدَقُ');

-- دین و زندگی (۳)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('دین و زندگی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس اول: هستی‌بخش'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس دوم: یگانه بی‌همتا'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس سوم: توحید و سبک زندگی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس چهارم: فقط برای او'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس پنجم: قدرت پرواز'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس ششم: سنت‌های خداوند در زندگی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس هفتم: بازگشت'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس هشتم: زندگی در دنیای امروز و عمل به احکام الهی'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس نهم: پایه‌های استوار'),
((SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'درس دهم: تمدن جدید و مسئولیت ما');

-- انگلیسی (۳)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('انگلیسی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Lesson 1: Sense of Appreciation'),
((SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Lesson 2: Look it Up!'),
((SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Lesson 3: Renewable Energy');

-- نگارش (۳)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('نگارش (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس اول: نوشته‌های ذهنی'),
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس دوم: نوشته‌های علمی'),
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس سوم: مستندنگاری'),
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس چهارم: نقد'),
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس پنجم: متن‌های عمومی'),
((SELECT id FROM books WHERE title = 'نگارش (۳)'), 'درس ششم: خلاصه‌نویسی');

-- سلامت و بهداشت
INSERT INTO books (title, inherent_grade_level_id) VALUES ('سلامت و بهداشت', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل اول: سلامت'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل دوم: تغذیه سالم و بهداشت مواد غذایی'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل سوم: پیشگیری از بیماری‌ها'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل چهارم: بهداشت روان'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل پنجم: پیشگیری از رفتارهای پرخطر'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل ششم: محیط کار و زندگی سالم'),
((SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'فصل هفتم: پیشگیری از حوادث خانگی و اصول ایمنی و کمک‌های اولیه');

-- مدیریت خانواده و سبک زندگی
INSERT INTO books (title, inherent_grade_level_id) VALUES ('مدیریت خانواده و سبک زندگی', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس اول: خانه و خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس دوم: رشدو تکامل انسان'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس سوم: مراحل تشکیل خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس چهارم: کارکردهای خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس پنجم: مدیریت منابع در خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس ششم: سلامت در خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس هفتم: خانواده و انواع آن'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس هشتم: آسیب شناسی خانواده'),
((SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'درس نهم: مدیریت خانواده');


-- Step 6.2: Seed Books & Lessons for 12th Grade (Specialized - تخصصی)

-- شیمی (۳) (مشترک ریاضی و تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('شیمی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'شیمی (۳)'), 'فصل اول: مولکول‌ها در خدمت تندرستی'),
((SELECT id FROM books WHERE title = 'شیمی (۳)'), 'فصل دوم: آسایش و رفاه در سایه شیمی'),
((SELECT id FROM books WHERE title = 'شیمی (۳)'), 'فصل سوم: شیمی جلوه‌ای از هنر، زیبایی و ماندگاری'),
((SELECT id FROM books WHERE title = 'شیمی (۳)'), 'فصل چهارم: شیمی، راهی به سوی آینده‌ای روشن‌تر');

-- حسابان (۲) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('حسابان (۲)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'حسابان (۲)'), 'فصل اول: تابع'),
((SELECT id FROM books WHERE title = 'حسابان (۲)'), 'فصل دوم: مثلثات'),
((SELECT id FROM books WHERE title = 'حسابان (۲)'), 'فصل سوم: حدهای نامتناهی - حد در بی‌نهایت'),
((SELECT id FROM books WHERE title = 'حسابان (۲)'), 'فصل چهارم: مشتق'),
((SELECT id FROM books WHERE title = 'حسابان (۲)'), 'فصل پنجم: کاربردهای مشتق');

-- هندسه (۳) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('هندسه (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'هندسه (۳)'), 'فصل اول: ماتریس و کاربردها'),
((SELECT id FROM books WHERE title = 'هندسه (۳)'), 'فصل دوم: مقاطع مخروطی'),
((SELECT id FROM books WHERE title = 'هندسه (۳)'), 'فصل سوم: بردارها');

-- ریاضیات گسسته (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ریاضیات گسسته', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ریاضیات گسسته'), 'فصل اول: آشنایی با نظریۀ اعداد'),
((SELECT id FROM books WHERE title = 'ریاضیات گسسته'), 'فصل دوم: گراف و مدل‌سازی'),
((SELECT id FROM books WHERE title = 'ریاضیات گسسته'), 'فصل سوم: ترکیبیات (شمارش)');

-- فیزیک (۳) (ریاضی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۳) - ریاضی', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل اول: حرکت بر خط راست'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل دوم: دینامیک و حرکت دایره‌ای'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل سوم: نوسان و موج'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل چهارم: برهم‌کنش‌های موج'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل پنجم: فیزیک اتمی'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'فصل ششم: فیزیک هسته‌ای');

-- ریاضی (۳) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ریاضی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل اول: تابع'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل دوم: مثلثات'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل سوم: حدهای نامتناهی و حد در بی‌نهایت'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل چهارم: مشتق'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل پنجم: کاربرد مشتق'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل ششم: هندسه'),
((SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'فصل هفتم: احتمال');

-- فیزیک (۳) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فیزیک (۳) - تجربی', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فیزیک (۳) - تجربی'), 'فصل اول: حرکت بر خط راست'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - تجربی'), 'فصل دوم: دینامیک'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - تجربی'), 'فصل سوم: نوسان و موج'),
((SELECT id FROM books WHERE title = 'فیزیک (۳) - تجربی'), 'فصل چهارم: فیزیک اتمی و هسته‌ای');

-- زیست شناسی (۳) (تجربی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('زیست شناسی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل اول: مولکول‌های اطلاعاتی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل دوم: جریان اطلاعات در یاخته'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل سوم: انتقال اطلاعات در نسل‌ها'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل چهارم: تغییر در اطلاعات وراثتی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل پنجم: از ماده به انرژی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل ششم: از انرژی به ماده'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل هفتم: فناوری‌های نوین زیستی'),
((SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'فصل هشتم: رفتارهای جانوران');

-- عربی، زبان تخصصی رشته انسانی (۳)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('عربی، زبان تخصصی رشته انسانی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۳)'), 'اَلدَّرْسُ الْأَوَّلُ: اَلْمُنادیٰ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۳)'), 'اَلدَّرْسُ الثّاني: اَلْمُسْتَثنیٰ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۳)'), 'اَلدَّرْسُ الثّالِثُ: اَلْاِسْتِفْهامُ وَ التَّعَجُّبُ'),
((SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۳)'), 'اَلدَّرْسُ الرّابِعُ: مَفْهومُ الْحَصْرِ وَ التَّقْدیمِ');

-- علوم و فنون ادبی (۳) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('علوم و فنون ادبی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس اول: تاریخ ادبیات دورۀ بیداری و مشروطه'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس دوم: اختیارات وزنی (۱)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس سوم: مراعات نظیر، تلمیح، تضمین'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس چهارم: سبک شناسی دورۀ بیداری و مشروطه'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس پنجم: اختیارات وزنی (۲)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس ششم: لف و نشر، تضاد، متناقض نما'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس هفتم: تاریخ ادبیات معاصر (نظم و نثر)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس هشتم: اختیارات شاعری'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس نهم: اغراق، ایهام، ایهام تناسب'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس دهم: سبک شناسی دورۀ معاصر (نظم و نثر)'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس یازدهم: وزن در شعر نیمایی'),
((SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'درس دوازدهم: حسن تعلیل، اسلوب معادله، حس آمیزی');

-- فلسفه (۲) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('فلسفه (۲)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس اول: هستی و چیستی'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس دوم: جهان ممکنات'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس سوم: جهان علّی و معلولی'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس چهارم: کدامین تصویر از جهان؟'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس پنجم: خدا در فلسفه (۱)'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس ششم: خدا در فلسفه (۲)'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس هفتم: عقل در فلسفه (۱)'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس هشتم: عقل در فلسفه (۲)'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس نهم: آغاز فلسفۀ جدید در اروپا'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس دهم: دورۀ جدید فلسفه در اروپا'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس یازدهم: فلسفۀ معاصر اروپایی'),
((SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'درس دوازدهم: حکمت و حکومت');

-- ریاضی و آمار (۳) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('ریاضی و آمار (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'ریاضی و آمار (۳)'), 'فصل اول: شمارش'),
((SELECT id FROM books WHERE title = 'ریاضی و آمار (۳)'), 'فصل دوم: احتمال'),
((SELECT id FROM books WHERE title = 'ریاضی و آمار (۳)'), 'فصل سوم: آمار');

-- جامعه شناسی (۳) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('جامعه شناسی (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس اول: ذخیرۀ دانشی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس دوم: علوم اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس سوم: جامعه‌شناسی تبیینی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس چهارم: جامعه‌شناسی تفسیری'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس پنجم: جامعه‌شناسی انتقادی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس ششم: کنش اجتماعی و نظام اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس هفتم: ساختار اجتماعی'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس هشتم: تحولات ساختاری'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس نهم: بازاندیشی دربارۀ علم'),
((SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'درس دهم: افق‌ها و چالش‌های پیش روی علوم اجتماعی');

-- تاریخ (۳) (انسانی)
INSERT INTO books (title, inherent_grade_level_id) VALUES ('تاریخ (۳)', (SELECT id FROM grades WHERE name = 'دوازدهم'));
INSERT INTO lessons (book_id, name) VALUES
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس اول: تاریخ‌نگاری و تحولات آن در دورۀ معاصر'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس دوم: ایران و جهان در آستانۀ دورۀ معاصر'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس سوم: سیاست و حکومت در عصر قاجار'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس چهارم: اوضاع اجتماعی، اقتصادی و فرهنگی عصر قاجار'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس پنجم: نهضت مشروطۀ ایران'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس ششم: جنگ جهانی اول و ایران'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس هفتم: ایران در دورۀ حکومت رضاشاه'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس هشتم: جنگ جهانی دوم و جهانِ پس از آن'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس نهم: نهضت ملی شدن صنعت نفت ایران'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس دهم: انقلاب اسلامی'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس یازدهم: استقرار و تثبیت نظام جمهوری اسلامی'),
((SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'درس دوازدهم: جنگ تحمیلی و دفاع مقدس');


-- =================================================================
--                              لینک دهی کتاب ها
-- =================================================================

-- Step 7: Seed book_roles to link books to curriculum

-- ** عمومی - پایه دهم **
-- فارسی (۱)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فارسی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فارسی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'فارسی (۱)'), 'Core');
-- دین و زندگی (۱)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 'Core');
-- انگلیسی (۱)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 'Core');
-- نگارش (۱)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'نگارش (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'نگارش (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'نگارش (۱)'), 'Core');
-- آمادگی دفاعی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 'Core');

-- ** تخصصی - پایه دهم **
-- ریاضی و تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'شیمی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'شیمی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'ریاضی (۱)'), 'Core');
-- ریاضی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فیزیک (۱) - ریاضی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'هندسه (۱)'), 'Core');
-- تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فیزیک (۱) - تجربی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'زیست شناسی (۱)'), 'Core');
-- انسانی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'منطق'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'اقتصاد'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'تاریخ (۱)'), 'Core');

-- ** عمومی - پایه یازدهم **
-- فارسی (۲)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فارسی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فارسی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'فارسی (۲)'), 'Core');
-- دین و زندگی (۲)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۲)'), 'Core');
-- انگلیسی (۲)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'انگلیسی (۲)'), 'Core');
-- نگارش (۲)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'نگارش (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'نگارش (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'نگارش (۲)'), 'Core');

-- ** تخصصی - پایه یازدهم **
-- ریاضی و تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'شیمی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'شیمی (۲)'), 'Core');
-- ریاضی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'حسابان (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'آمار و احتمال'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'هندسه (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فیزیک (۲) - ریاضی'), 'Core');
-- تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'ریاضی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فیزیک (۲) - تجربی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'زیست شناسی (۲)'), 'Core');
-- انسانی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'فلسفه (۱)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'روان شناسی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'جامعه شناسی (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'یازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'تاریخ (۲)'), 'Core');

-- ** عمومی - پایه دوازدهم **
-- فارسی (۳)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فارسی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فارسی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'فارسی (۳)'), 'Core');
-- دین و زندگی (۳)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'دین و زندگی (۳)'), 'Core');
-- انگلیسی (۳)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'انگلیسی (۳)'), 'Core');
-- نگارش (۳)
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'نگارش (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'نگارش (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'نگارش (۳)'), 'Core');
-- سلامت و بهداشت
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'سلامت و بهداشت'), 'Core');
-- مدیریت خانواده و سبک زندگی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'مدیریت خانواده و سبک زندگی'), 'Core');


-- ** تخصصی - پایه دوازدهم **
-- ریاضی و تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'عربی، زبان قرآن (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'شیمی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'شیمی (۳)'), 'Core');
-- ریاضی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'حسابان (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'هندسه (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'ریاضیات گسسته'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'ریاضی فیزیک'), (SELECT id FROM books WHERE title = 'فیزیک (۳) - ریاضی'), 'Core');
-- تجربی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'ریاضی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'فیزیک (۳) - تجربی'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم تجربی'), (SELECT id FROM books WHERE title = 'زیست شناسی (۳)'), 'Core');
-- انسانی
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'فلسفه (۲)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'ریاضی و آمار (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'جامعه شناسی (۳)'), 'Core');
INSERT INTO book_roles (target_student_grade_id, major_id, book_id, role) VALUES ((SELECT id FROM grades WHERE name = 'دوازدهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), (SELECT id FROM books WHERE title = 'تاریخ (۳)'), 'Core');


-- Step 8: Insert schedule template (Corrected to use Grade 10, Major 3)
INSERT INTO schedule_templates (name, target_grade_id, target_major_id, total_study_blocks_per_week) VALUES
('دهم انسانی - ۲۴ بلوک', (SELECT id FROM grades WHERE name = 'دهم'), (SELECT id FROM majors WHERE name = 'علوم انسانی'), 24);


-- Insert sample schedule templates
INSERT INTO schedule_templates (name, target_grade_id, target_major_id, total_study_blocks_per_week) VALUES
('Light Schedule', 1, 1, 12),
('Standard Schedule', 1, 1, 18), 
('Intensive Schedule', 1, 1, 24);

-- Insert template rules for books
INSERT INTO template_rules (template_id, book_id, default_frequency, priority_slot, time_preference, consecutive_sessions) VALUES
(1, 1, 4, 'first', 'morning', true),
(1, 2, 4, NULL, 'afternoon', false),
(1, 3, 4, NULL, NULL, false),
(2, 1, 6, 'first', 'morning', true),
(2, 2, 6, NULL, 'afternoon', false), 
(2, 3, 6, NULL, NULL, false),
(3, 1, 8, 'first', 'morning', true),
(3, 2, 8, NULL, 'afternoon', false),
(3, 3, 8, NULL, NULL, false);


-- Step 9: Template rules (Corrected book titles and frequencies to sum to 24 for Grade 10 Humanities)
INSERT INTO template_rules (template_id, book_id, default_frequency, scheduling_hints, consecutive_sessions, time_preference, priority_slot) VALUES
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'فارسی (۱)'), 3, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'عربی، زبان تخصصی رشته انسانی (۱)'), 3, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'دین و زندگی (۱)'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'انگلیسی (۱)'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'نگارش (۱)'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'آمادگی دفاعی'), 1, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'علوم و فنون ادبی (۱)'), 3, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'منطق'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'اقتصاد'), 2, NULL, FALSE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'جامعه شناسی (۱)'), 2, 'contiguous_pair', TRUE, NULL, NULL),
((SELECT id FROM schedule_templates WHERE name = 'دهم انسانی - ۲۴ بلوک'), (SELECT id FROM books WHERE title = 'تاریخ (۱)'), 2, NULL, FALSE, NULL, NULL);

