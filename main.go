// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	cors2 "github.com/sagar510/student-go/internal/cors"
	"github.com/sagar510/student-go/internal/repository"
	controller "github.com/sagar510/student-go/internal/services"
)

func main() {

	repository.InitDB()

	db := repository.GetDB()

	//All Repositories
	teachCourseRepository := repository.NewCourseRepository(db)
	enrolCourseRepository := repository.NewStudentCourseRepository(db)

	courseHandler := controller.NewCourseController(teachCourseRepository)
	enrolHandler := controller.NewStudentCourseController(enrolCourseRepository)

	// Creating a Gorilla Mux router
	router := mux.NewRouter()

	// Using the Authorization Middleware for all routes
	//router.Use(loggingMiddleware)

	router.HandleFunc("/createcourse", courseHandler.CreateCourse).Methods("POST")
	router.HandleFunc("/teachcourse", courseHandler.TeachCourse).Methods("POST")

	router.HandleFunc("/viewcoursesteacher", courseHandler.ViewCourses).Methods("GET")
	router.HandleFunc("/viewcoursesstudent", enrolHandler.ViewCourses).Methods("GET")
	router.HandleFunc("/enrol", enrolHandler.EnrolCourse).Methods("POST")

	corsHandler := cors2.Mangement(router)

	// Wrapping the router with the CORS handler
	handler := corsHandler.Handler(router)

	// Starting the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", handler))

}
