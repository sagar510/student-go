package repository

import (
	"fmt"

	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {

	username := "root"
	password := ""
	host := "localhost"
	databaseName := "studentmarks_development"

	//dsn := "root:@tcp(localhost:3306/studentmarks_development?charset=utf8mb4&parseTime=True&loc=Local)"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", username, password, host, databaseName)

	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableAutomaticSoftDelete: true,
	})

	if err != nil {
		panic("failed to connect to database")

	}

	db.AutoMigrate(&domain.Students{}, &domain.Courses{}, &domain.Enrolments{}, &domain.Teacher_courses{}, &domain.Teachers{})

}

func GetDB() *gorm.DB {
	return db
}
