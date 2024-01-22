package repository

import (
	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/gorm"
)

type CourseRepository interface {
	CreateCourse(course *domain.Courses) error
	TeachCourse(course_id uint, tchr_id uint) error
}

//Implementation of above interface starts here

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) CreateCourse(course *domain.Courses) error {
	/*course.Created_at = time.Now()
	course.Updated_at = time.Now()*/

	course.Ongoing = true
	return r.db.Create(course).Error
}

func (r *courseRepository) TeachCourse(course_id uint, tchr_id uint) error {
	var teacher_course domain.Teacher_courses
	//var teacher domain.Teachers
	//var course domain.Courses

	/*course.Created_at = time.Now()
	course.Updated_at = time.Now()

	// Attempt to fetch the teacher_course record from the database by ID
	if err := r.db.Unscoped().First(&course, course_id).Error; err != nil {
		// If there's an error during the database query, return the error
		return err
	}

	if err := r.db.Unscoped().First(&teacher, tchr_id).Error; err != nil {
		return err
	}
	*/
	if errr := r.db.Where("course_id = ? AND teacher_id = ?", course_id, tchr_id).First(&teacher_course).Error; errr == nil {
		return errr
	}

	teacher_course.Course_id = course_id
	teacher_course.Teacher_id = tchr_id

	return r.db.Create(&teacher_course).Error
}
