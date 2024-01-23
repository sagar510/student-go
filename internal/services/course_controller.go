package controller

import (
	"encoding/json"
	"net/http"

	"github.com/sagar510/student-go/internal/domain"
	"github.com/sagar510/student-go/internal/repository"
	"github.com/sagar510/student-go/internal/utils"
)

/*type CourseControllerInterface interface{
	func (cc *courseController) CreateCourse(w http.ResponseWriter,r *http.Request);
}*/

//Implementation starts here

type CourseController struct {
	courseRepository repository.CourseRepository
}

func NewCourseController(courseRepository repository.CourseRepository) *CourseController {
	return &CourseController{courseRepository}
}

func (cc *CourseController) CreateCourse(w http.ResponseWriter, r *http.Request) {
	var course domain.Courses

	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = cc.courseRepository.CreateCourse(&course)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//utils.RespondJSON(w, course)
}

func (cc *CourseController) TeachCourse(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)

	courseID, err := strconv.Atoi(vars["courseID"])
	if err != nil {
		http.Error(w, "Invalid Course ID", http.StatusBadRequest)
		return
	}

	teacherID, err := strconv.Atoi(vars["teacherID"])
	if err != nil {
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}*/

	/*var teachercourse domain.Teacher_courses

	err := json.NewDecoder(r.Body).Decode(&teachercourse)
	*/

	// Parse JSON from the request body
	var requestData map[string]int
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Extract courseID and teacherID from the parsed JSON
	courseID, ok1 := requestData["courseID"]
	teacherID, ok2 := requestData["teacherID"]
	if !ok1 || !ok2 {
		http.Error(w, "Invalid Course or Teacher ID in JSON", http.StatusBadRequest)
		return
	}

	err := cc.courseRepository.TeachCourse(uint(courseID), uint(teacherID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func (cc *CourseController) ViewCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := cc.courseRepository.ViewCourses()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondJSON(w, courses)
}

func (cc *CourseController) ATeacherCourse(w http.ResponseWriter, r *http.Response) {
	var requestData map[string]int
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	teacherID, ok := requestData["teacherID"]

	if !ok {
		http.Error(w, "Teacher ID in JSON", http.StatusBadRequest)
		return
	}

	list, err := cc.courseRepository.ATeacherCourse(uint(teacherID))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	utils.RespondJSON(w, list)
}
