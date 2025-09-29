package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func PermissionIndex(c *gin.Context) {
	var permissions []models.Permission

	if err := config.DB.Find(&permissions).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": permissions, "message": nil})
}

func PermissionStore(c *gin.Context) {
	var data []models.Permission

	// Parse request JSON
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"status":  false,
			"data":    nil,
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	// Validasi sederhana
	for _, perm := range data {
		if strings.TrimSpace(perm.Name) == "" {
			c.JSON(400, gin.H{
				"status":  false,
				"data":    nil,
				"message": "permission name cannot be empty",
			})
			return
		}
	}

	// Pakai transaction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Upsert: kalau name sudah ada → update parent_menu
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},                                 // kolom unik
			DoUpdates: clause.AssignmentColumns([]string{"parent_menu", "updated_at"}), // update kolom tertentu
		}).Create(&data).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{
			"status":  false,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"status":  true,
		"data":    data,
		"message": "success",
	})
}
