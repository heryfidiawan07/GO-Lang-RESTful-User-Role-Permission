package config

import (
	"app/models"
	"fmt"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	gotenv.Load()

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:@tcp(127.0.0.1:3306)/go_rest_role_permissions?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := os.Getenv("DATABASE")
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	var err error
	p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect database !" + err.Error())
	}

	// defer db.Close()

	fmt.Println("* * * * * * * * * ")
	fmt.Println("Database Connected")
	fmt.Println("* * * * * * * * * ")

	// Migrate Schema
	DB.AutoMigrate(&models.Role{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Permission{})
	DB.AutoMigrate(&models.RolePermission{})
	DB.AutoMigrate(&models.RefreshToken{})
}
