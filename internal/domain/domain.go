package domain

import (
	"math/big"

	"gorm.io/gorm"
)

type Students struct {
	id           big.Int
	Student_name string
	Phone        string
	Address      string

	user_id  big.Int
	admin_id big.Int

	enrolments []Enrolments
}

type Courses struct {
	gorm.Model
	//Id          uint
	Course_name string
	Ongoing     bool
	Duration    int
	//Created_at  time.Time `gorm:"-"`
	//Updated_at  time.Time `gorm:"-"`

	enrolments      []Enrolments
	teacher_courses []Teacher_courses
}

type Enrolments struct {
	id    big.Int
	Marks int
	Year  int

	student_id big.Int
	course_id  big.Int
}

type Teachers struct {
	id           big.Int
	Teacher_name string
	Phone        string
	Address      string

	user_id  big.Int
	admin_id big.Int

	teacher_courses []Teacher_courses
}

type Teacher_courses struct {
	gorm.Model //`gorm:"soft_delete: false"`

	//id   uint
	Year int

	Teacher_id uint
	Course_id  uint

	enrolments      []Enrolments
	teacher_courses []Teacher_courses
}
