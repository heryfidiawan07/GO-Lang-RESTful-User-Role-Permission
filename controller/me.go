package controller

import (
	// "fmt"
	"restfull-api/config"
	"restfull-api/models"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	var user models.User
	
	if err := config.DB.First(&user, "id = ?", c.MustGet("jwt_user_id")).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	var role models.Role
	config.DB.First(&role, "id = ?", user.RoleId)

	var rolePermissions []models.RolePermission
	config.DB.Where("role_id = ?", user.RoleId).Find(&rolePermissions)

	rolePermissionId := make([]string, len(rolePermissions))
	for key,value := range rolePermissions {
		rolePermissionId[key] = value.PermissionId
    }
	
	var permissions []models.Permission
	config.DB.Where("id IN ?", rolePermissionId).Find(&permissions)

	actions := make([]string, len(permissions))
	for key,value := range permissions {
		actions[key] = value.Name
    }

	temp := struct {
		models.Role
		Permissions []models.Permission
	}{role, permissions}

	data := struct {
		models.User
		Role models.RolePermissionCombine
		Actions []string
	}{user, temp, actions}

	c.JSON(200, gin.H{"status": true, "data": data, "message": nil})
}
