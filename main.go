// main.go
package main

import (
	"fmt"

	"github.com/sagar510/student-go/internal/repository"
)

func main() {

	repository.InitDB()

	db := repository.GetDB()

	repository.NewCourseRepository(db)

	fmt.Println("Hello, World!")
}
