package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"encoding/json"
	"io/ioutil"

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
	body, err := ioutil.ReadAll(c.Request.Body)
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
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(201, gin.H{"status": true, "data": data, "message": "Data created successfully"})
}
