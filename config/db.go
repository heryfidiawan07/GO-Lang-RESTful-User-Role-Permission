package config

import (
	"fmt"
	"restfull-api/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/sk_restfull_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if(err != nil) {
		panic("Failed connect database !" + err.Error())
	}

	// defer db.Close()

	fmt.Println("* * * * * * * * * ")
	fmt.Println("Database Connected")
	fmt.Println("* * * * * * * * * ")

	// Migrate Schema
	db.AutoMigrate(&models.Role{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Permission{})
	db.AutoMigrate(&models.RolePermission{})

	DB = db
}