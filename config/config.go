package config

import (
	admins "project/model/admin"
	"project/model/books"
	"project/model/loan"
	"project/model/status"
	"project/model/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "root:adarizki123@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "admin:adarizki123@tcp(db-book-library.cpzrb8hxmi0a.us-east-2.rds.amazonaws.com:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}
func InitDBTest() {
	// dsn := "root:adarizki123@tcp(127.0.0.1:3306)/library_test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "admin:adarizki123@tcp(db-book-library.cpzrb8hxmi0a.us-east-2.rds.amazonaws.com:3306)/library_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	InitMigrateTest()
}

func InitMigrate() {
	DB.AutoMigrate(&users.User{})
	DB.AutoMigrate(&books.Book{})
	DB.AutoMigrate(&loan.Loan{})
	DB.AutoMigrate(&admins.Admin{})
	DB.AutoMigrate(&status.Status{})
}
func InitMigrateTest() {
	DB.Migrator().DropTable(&users.User{})
	DB.AutoMigrate(&users.User{})
	DB.Migrator().DropTable(&admins.Admin{})
	DB.AutoMigrate(&admins.Admin{})
	DB.Migrator().DropTable(&books.Book{})
	DB.AutoMigrate(&books.Book{})
	DB.Migrator().DropTable(&status.Status{})
	DB.AutoMigrate(&status.Status{})
}
