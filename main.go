// main.go
package main

import (
	"fmt"

	"github.com/sagar510/student-go/internal/repository"
)

func main() {

	repository.InitDB()

	fmt.Println("Hello, World!")
}
