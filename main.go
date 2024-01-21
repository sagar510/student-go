// main.go
package main

import (
	"fmt"

	"github.com/sagar510/student-go/internal/repository"
	"github.com/sagar510/student-go/internal/service"
)

func main() {

	repository.InitDB()

	db := repository.GetDB()

	repo := repository.NewCourseRepository(db)

	service.NewCourseService(repo)

	fmt.Println("Hello, World!")
}
