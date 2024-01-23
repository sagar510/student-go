package repository

import (
	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *domain.Courses) error
	TeachCourse(course_id uint, tchr_id uint) error
	ViewCourses() ([]domain.Courses, error)
	ATeacherCourse(teacher_id uint) ([]domain.Teacher_courses, error)
}

//Implementation of above interface starts here

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) CreateCourse(course *domain.Courses) error {
	course.Ongoing = true
	return r.db.Create(course).Error
}

func (r *courseRepository) TeachCourse(course_id uint, tchr_id uint) error {
	var teacher_course domain.Teacher_courses

	if errr := r.db.Where("course_id = ? AND teacher_id = ?", course_id, tchr_id).First(&teacher_course).Error; errr == nil {
		return errr
	}

	teacher_course.Course_id = course_id
	teacher_course.Teacher_id = tchr_id

	return r.db.Create(&teacher_course).Error
	//r.db.Create(&student).Error
}

func (r *courseRepository) ViewCourses() ([]domain.Courses, error) {
	var courses []domain.Courses

	if err := r.db.Where("ongoing=?", true).Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepository) ATeacherCourse(teacher_id uint) ([]domain.Teacher_courses, error) {
	var list []domain.Teacher_courses

	if err := r.db.Where("teacher_id=?", teacher_id).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
