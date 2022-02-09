package controller

import (
	// "fmt"
	"restfull-api/config"
	"restfull-api/models"
	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": users, "message": nil})
}

func UserCreate(c *gin.Context) {
	data := models.User {
		Username: c.PostForm("username"),
		Name: c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "Data created successfully"})
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	data := models.User{
		Username: c.PostForm("username"),
		Name: c.PostForm("name"),
		Email: c.PostForm("email"),
	}

	if err := config.DB.Model(&user).Updates(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": user, "message": "Data updated successfully"})
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": user, "message": nil})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": nil, "message": "Data deleted successfully"})
}
