package repository

import (
	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/gorm"
)

type StudentCourseRepository interface {
	EnrolCourse(course_id uint, stu_id uint) error
	ViewCourses() ([]domain.Courses, error)
	//Enrol(teacher_id uint) ([]domain.Enrolments, error)
}

type studentcourseRepository struct {
	db *gorm.DB
}

func NewStudentCourseRepository(db *gorm.DB) StudentCourseRepository {
	return &studentcourseRepository{db}
}

func (s *studentcourseRepository) ViewCourses() ([]domain.Courses, error) {
	var courses []domain.Courses

	if err := s.db.Where("ongoing=?", true).Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *studentcourseRepository) EnrolCourse(course_id uint, stu_id uint) error {
	var enrol domain.Enrolments

	if err := r.db.Where("course_id = ? AND student_id = ?", course_id, stu_id).First(&enrol.Marks).Error; err == nil {
		return err
	}

	enrol.Course_id = course_id
	enrol.Student_id = stu_id

	return r.db.Create(&enrol).Error
	//r.db.Create(&student).Error
}

/*
func (r *studentcourseRepository) Enrol(teacher_id uint) ([]domain.Enrolments, error) {
	var list []domain.Enrolments

	if err := r.db.Where("teacher_id=?", teacher_id).Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
*/
