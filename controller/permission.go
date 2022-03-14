package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"github.com/gin-gonic/gin"
)

func PermissionIndex(c *gin.Context) {
	var permissions []models.Permission

	if err := config.DB.Find(&permissions).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": permissions, "message": nil})
}

func PermissionStore(c *gin.Context) {
	data := models.Permission {
		ParentMenu: c.PostForm("parent_menu"),
		ParentId: c.PostForm("parent_id"),
		Name: c.PostForm("name"),
		Alias: c.PostForm("alias"),
		Url: c.PostForm("url"),
		Icon: c.PostForm("icon"),
	}

	if err := config.DB.Create(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "Data created successfully"})
}