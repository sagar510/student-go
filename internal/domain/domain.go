package domain

import (
	"math/big"
)

type Students struct {
	id           big.Int
	student_name string
	phone        string
	address      string

	user_id  big.Int
	admin_id big.Int

	enrolments []Enrolments
}

type Courses struct {
	id          big.Int
	course_name string
	ongoing     bool
	duration    int

	enrolments      []Enrolments
	teacher_courses []Teacher_courses
}

type Enrolments struct {
	id    big.Int
	marks int
	year  int

	student_id big.Int
	course_id  big.Int
}

type Teachers struct {
	id           big.Int
	teacher_name string
	phone        string
	address      string

	user_id  big.Int
	admin_id big.Int

	teacher_courses []Teacher_courses
}

type Teacher_courses struct {
	id   big.Int
	year int

	teacher_id big.Int
	course_id  big.Int
}
