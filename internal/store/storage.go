package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrorNotFound       = errors.New("resource not found")
	ErrorDuplicateEmail = errors.New("duplicate email")
)

type Storage struct {
	Students           StudentStore
	Grades             GradeStore
	Majors             MajorStore
	Books              BookStore
	Lessons            LessonStore
	UnavailableTimes   UnavailableTimeStore
	WeeklyPlans        WeeklyPlanStore
	SubjectFrequencies SubjectFrequencyStore
	DailyPlans         DailyPlanStore
	StudySessions      StudySessionStore
	SessionReports     SessionReportStore
	ExamSchedules      ExamScheduleStore
	ExamScopeItems     ExamScopeItemStore
	ScheduleTemplates  ScheduleTemplateStore
	TemplateRules      TemplateRuleStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students:           &StudentModel{DB: db},
		Grades:             &GradeModel{DB: db},
		Majors:             &MajorModel{DB: db},
		Books:              &BookModel{DB: db},
		Lessons:            &LessonModel{DB: db},
		UnavailableTimes:   &UnavailableTimeModel{DB: db},
		WeeklyPlans:        &WeeklyPlanModel{DB: db},
		SubjectFrequencies: &SubjectFrequencyModel{DB: db},
		DailyPlans:         &DailyPlanModel{DB: db},
		StudySessions:      &StudySessionModel{DB: db},
		SessionReports:     &SessionReportModel{DB: db},
		ExamSchedules:      &ExamScheduleModel{DB: db},
		ExamScopeItems:     &ExamScopeItemModel{DB: db},
		ScheduleTemplates:  &ScheduleTemplateModel{DB: db},
		TemplateRules:      &TemplateRuleModel{DB: db},
	}
}

type StudentStore interface {
	Insert(ctx context.Context, student *Student) error
	Get(ctx context.Context, id int64) (*Student, error)
	Update(ctx context.Context, student *Student) error
	Delete(ctx context.Context, id int64) error
}

type GradeStore interface {
	Get(ctx context.Context, id int64) (*Grade, error)
	GetAll(ctx context.Context) ([]*Grade, error)
}

type MajorStore interface {
	Get(ctx context.Context, id int64) (*Major, error)
	GetAll(ctx context.Context) ([]*Major, error)
}

type BookStore interface {
	Get(ctx context.Context, id int64) (*Book, error)
	GetAllForCurriculum(ctx context.Context, gradeID, majorID int64) ([]*Book, error)
}

type LessonStore interface {
	Get(ctx context.Context, id int64) (*Lesson, error)
	GetAllForBook(ctx context.Context, bookID int64) ([]*Lesson, error)
}

type UnavailableTimeStore interface {
	Insert(ctx context.Context, ut *UnavailableTime) error
	Get(ctx context.Context, id int64) (*UnavailableTime, error)
	GetAllForStudent(ctx context.Context, studentID int64) ([]*UnavailableTime, error)
	Update(ctx context.Context, ut *UnavailableTime) error
	Delete(ctx context.Context, id int64) error
}

type WeeklyPlanStore interface {
	Insert(ctx context.Context, wp *WeeklyPlan) error
	Get(ctx context.Context, id int64) (*WeeklyPlan, error)
	GetAllForStudent(ctx context.Context, studentID int64) ([]*WeeklyPlan, error)
	Update(ctx context.Context, wp *WeeklyPlan) error
	Delete(ctx context.Context, id int64) error
}

type SubjectFrequencyStore interface {
	Insert(ctx context.Context, sf *SubjectFrequency) error
	GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*SubjectFrequency, error)
	Update(ctx context.Context, sf *SubjectFrequency) error
	Delete(ctx context.Context, id int64) error
}

type DailyPlanStore interface {
	Insert(ctx context.Context, dp *DailyPlan) error
	Get(ctx context.Context, id int64) (*DailyPlan, error)
	GetByWeeklyPlanAndDate(ctx context.Context, weeklyPlanID int64, planDate time.Time) (*DailyPlan, error)
	GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*DailyPlan, error)
	Delete(ctx context.Context, id int64) error
}

type StudySessionStore interface {
	Insert(ctx context.Context, ss *StudySession) error
	Get(ctx context.Context, id int64) (*StudySession, error)
	GetAllForDailyPlan(ctx context.Context, dailyPlanID int64) ([]*StudySession, error)
	Update(ctx context.Context, ss *StudySession) error
	Delete(ctx context.Context, id int64) error
}

type SessionReportStore interface {
	Insert(ctx context.Context, sr *SessionReport) error
	GetForStudySession(ctx context.Context, studySessionID int64) (*SessionReport, error)
	Update(ctx context.Context, sr *SessionReport) error
	Delete(ctx context.Context, id int64) error
}

type ExamScheduleStore interface {
	Insert(ctx context.Context, es *ExamSchedule) error
	Get(ctx context.Context, id int64) (*ExamSchedule, error)
	GetAllForStudentCurriculum(ctx context.Context, gradeID, majorID int64) ([]*ExamSchedule, error)
	Update(ctx context.Context, es *ExamSchedule) error
	Delete(ctx context.Context, id int64) error
}

type ExamScopeItemStore interface {
	Insert(ctx context.Context, esi *ExamScopeItem) error
	GetAllForExam(ctx context.Context, examID int64) ([]*ExamScopeItem, error)
	Delete(ctx context.Context, id int64) error
}

type ScheduleTemplateStore interface {
	Get(ctx context.Context, id int64) (*ScheduleTemplate, error)
	GetAll(ctx context.Context, gradeID, majorID int64) ([]*ScheduleTemplate, error)
}

type TemplateRuleStore interface {
	Insert(ctx context.Context, tr *TemplateRule) error
	Get(ctx context.Context, id int64) (*TemplateRule, error)
	GetAllForTemplate(ctx context.Context, templateID int64) ([]*TemplateRule, error)
	Update(ctx context.Context, tr *TemplateRule) error
	Delete(ctx context.Context, id int64) error
}
