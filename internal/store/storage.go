// internal/store/store.go
package store

import (
	"context"
	"database/sql"
	"errors"
)

// Global errors
var (
	ErrorNotFound       = errors.New("resource not found")
	ErrorDuplicateEmail = errors.New("duplicate email")
)

// Storage is the main struct that holds all our data access types (interfaces).

type Storage struct {
	Students StudentStore
	Grades   GradeStore
	Majors   MajorStore
	Books    BookStore
	// ... and so on for every interface
}

// NewStorage creates a new Storage instance with all the data stores initialized.
func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Students: &StudentModel{DB: db},
		Grades:   &GradeModel{DB: db},
		Majors:   &MajorModel{DB: db},
		// ... and so on
	}
}

// --- DATA STORE INTERFACES ---

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

// ... continue defining an interface for every model ...
