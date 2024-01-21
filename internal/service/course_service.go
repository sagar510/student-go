package service

import (
	"github.com/sagar510/student-go/internal/domain"
	"github.com/sagar510/student-go/internal/repository"
)

type CourseService interface {
	CreateCourse(course *domain.Courses) error
}

// Course Service Implementation starts here

type courseService struct {
	courseRepository repository.CourseRepository
}

func NewCourseService(courseRepository repository.CourseRepository) CourseService {
	return &courseService{courseRepository}
}

func (s *courseService) CreateCourse(course *domain.Courses) error {
	return s.courseRepository.CreateCourse(course)
}
