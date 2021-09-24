package config

import (
	books "project/model/Books"
	"project/model/loan"
	"project/model/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:adarizki123@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&books.Book{})
	DB.AutoMigrate(&loan.Loan{})
}
