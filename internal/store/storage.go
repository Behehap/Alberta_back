// internal/store/storage.go
package store

import (
	"context"
	"database/sql"
	"errors"
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
	UnavailableTimes   UnavailableTimeStore
	WeeklyPlans        WeeklyPlanStore
	SubjectFrequencies SubjectFrequencyStore
	DailyPlans         DailyPlanStore
	StudySessions      StudySessionStore
	SessionReports     SessionReportStore
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students:           &StudentModel{DB: db},
		Grades:             &GradeModel{DB: db},
		Majors:             &MajorModel{DB: db},
		Books:              &BookModel{DB: db},
		UnavailableTimes:   &UnavailableTimeModel{DB: db},
		WeeklyPlans:        &WeeklyPlanModel{DB: db},
		SubjectFrequencies: &SubjectFrequencyModel{DB: db},
		DailyPlans:         &DailyPlanModel{DB: db},
		StudySessions:      &StudySessionModel{DB: db},
		SessionReports:     &SessionReportModel{DB: db},
	}
}

// --- INTERFACES ---

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

type UnavailableTimeStore interface {
	Insert(ctx context.Context, ut *UnavailableTime) error
	GetAllForStudent(ctx context.Context, studentID int64) ([]*UnavailableTime, error)
}

type WeeklyPlanStore interface {
	Insert(ctx context.Context, wp *WeeklyPlan) error
	Get(ctx context.Context, id int64) (*WeeklyPlan, error)
	GetAllForStudent(ctx context.Context, studentID int64) ([]*WeeklyPlan, error)
}

type SubjectFrequencyStore interface {
	Insert(ctx context.Context, sf *SubjectFrequency) error
	GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*SubjectFrequency, error)
}

type DailyPlanStore interface {
	Insert(ctx context.Context, dp *DailyPlan) error
	GetAllForWeeklyPlan(ctx context.Context, weeklyPlanID int64) ([]*DailyPlan, error)
	Get(ctx context.Context, id int64) (*DailyPlan, error)
}

type StudySessionStore interface {
	Insert(ctx context.Context, ss *StudySession) error
	GetAllForDailyPlan(ctx context.Context, dailyPlanID int64) ([]*StudySession, error)
	Get(ctx context.Context, id int64) (*StudySession, error)
	Update(ctx context.Context, ss *StudySession) error
	Delete(ctx context.Context, id int64) error
}

type SessionReportStore interface {
	Insert(ctx context.Context, sr *SessionReport) error
	GetForStudySession(ctx context.Context, studySessionID int64) (*SessionReport, error)
}
