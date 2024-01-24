package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sagar510/student-go/internal/repository"
	"github.com/sagar510/student-go/internal/utils"
)

type StudentCourseController struct {
	studentcourseRepository repository.StudentCourseRepository
}

func NewStudentCourseController(studentcourseRepository repository.StudentCourseRepository) *StudentCourseController {
	return &StudentCourseController{studentcourseRepository}
}

func (cc *StudentCourseController) ViewCourses(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	isValid, role := ValidateToken(token)

	if isValid {

		if role == "student" {

			courses, err := cc.studentcourseRepository.ViewCourses()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			utils.RespondJSON(w, courses)
		} else {
			w.WriteHeader(http.StatusForbidden)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (cc *StudentCourseController) EnrolCourse(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]int
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	courseID, ok1 := requestData["courseID"]
	studentID, ok2 := requestData["studentID"]

	if !ok1 || !ok2 {
		http.Error(w, "Invalid Course or Student ID in JSON", http.StatusBadRequest)
		return
	}

	err := cc.studentcourseRepository.EnrolCourse(uint(courseID), uint(studentID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
