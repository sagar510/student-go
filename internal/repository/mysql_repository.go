package repository

import (
	"github.com/sagar510/student-go/internal/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:@tcp(localhost:3306/studentmarks_development?charset=utf8mb4&parseTime=True&loc=Local)"
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")

	}

	db.AutoMigrate(&domain.Students{})

}

func GetDB() *gorm.DB {
	return db
}
