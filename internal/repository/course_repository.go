package repository

import (
	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *domain.Courses) error
}

//Implementation of above interface starts here

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) CreateCourse(course *domain.Courses) error {
	return r.db.Create(course).Error
}
