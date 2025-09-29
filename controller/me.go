package controller

import (
	// "fmt"
	"app/config"
	"app/models"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	var user models.User

	if err := config.DB.Preload("Role.Permissions").First(&user, "id = ?", c.MustGet("jwt_user_id")).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	var actions []string
	if user.Role != nil {
		for _, p := range user.Role.Permissions {
			actions = append(actions, p.Name)
		}
	}

	result := struct {
		models.User
		Actions []string
	}{user, actions}

	c.JSON(200, gin.H{"status": true, "data": result, "message": nil})
}
