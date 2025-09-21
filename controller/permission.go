package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"encoding/json"
	"io"
	"strings"

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
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	var data []models.Permission
	if err := json.Unmarshal(body, &data); err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}

	if err := config.DB.Create(&data).Error; err != nil {
		// Check if it's a unique constraint violation
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			c.JSON(409, gin.H{"status": false, "data": nil, "message": "Permission name already exists"})
		} else {
			c.JSON(500, gin.H{"status": false, "data": nil, "message": err.Error()})
		}
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "Data created successfully"})
}
