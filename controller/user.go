package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"app/request"
	"app/helper"
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

func UserStore(c *gin.Context) {
	var valid request.UserStore
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	data := models.User {
		Username: valid.Username,
		Name: valid.Name,
		Email: valid.Email,
		Password: helper.HashAndSalt([]byte(valid.Password)),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "Data created successfully"})
}

func UserUpdate(c *gin.Context) {
	var valid request.UserUpdate
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	id := c.Param("id")
	var user models.User

	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	data := models.User{
		Username: valid.Username,
		Name: valid.Name,
		Email: valid.Email,
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


func ChangePassword(c *gin.Context) {
	var valid request.ChangePassword
	if err := c.ShouldBind(&valid); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", c.MustGet("jwt_user_id")).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	validPassword := helper.ComparePassword(user.Password, []byte(valid.OldPassword))
	if !validPassword {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Old password invalid !"})
		return
	}

	data := models.User{
		Password: helper.HashAndSalt([]byte(valid.NewPassword)),
	}

	if err := config.DB.Model(&user).Updates(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": user, "message": "Change password successfully"})
}