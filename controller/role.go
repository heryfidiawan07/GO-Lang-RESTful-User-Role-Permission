package controller

import (
	// "fmt"
	"restfull-api/config"
	"restfull-api/models"
	"github.com/gin-gonic/gin"
)

func RoleIndex(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": roles, "message": nil})
}

func RoleCreate(c *gin.Context) {
	role := models.Role {
		Name: c.PostForm("name"),
	}

	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}
	
	for _, value := range c.PostFormMap("permissions") {
		permissions := models.RolePermission{
			RoleId: role.Id,
			PermissionId: value,
		}
		if err := config.DB.Create(&permissions).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
			return
		}
	}

	c.JSON(201, gin.H{"status": true, "data": role, "message": "Data created successfully"})
}

func RoleUpdate(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if err := config.DB.First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	data := models.Role{
		Name: c.PostForm("name"),
	}

	if err := config.DB.Model(&role).Updates(&data).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	if err := config.DB.Delete(models.RolePermission{}, "role_id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	for _, value := range c.PostFormMap("permissions") {
		permissions := models.RolePermission{
			RoleId: role.Id,
			PermissionId: value,
		}
		if err := config.DB.Create(&permissions).Error; err != nil {
			c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
			return
		}
	}

	c.JSON(200, gin.H{"status": true, "data": role, "message": "Data updated successfully"})
}

func RoleShow(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if err := config.DB.First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	var rolePermissions []models.RolePermission
	config.DB.Where("role_id = ?", id).Find(&rolePermissions)

	rolePermissionId := make([]string, len(rolePermissions))
	for key,value := range rolePermissions {
		rolePermissionId[key] = value.PermissionId
    }
	
	var permissions []models.Permission
	config.DB.Where("id IN ?", rolePermissionId).Find(&permissions)

	// data := models.RolePermissionCombine{
	// 	Role: role,
	// 	Permissions: permissions,
	// }
	data := struct {
		models.Role
		Permissions []models.Permission
	}{role, permissions}

	c.JSON(200, gin.H{"status": true, "data": data, "permissions": permissions, "message": nil})
}

func RoleDelete(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if err := config.DB.First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found !"})
		return
	}

	if err := config.DB.Delete(&role).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	if err := config.DB.Delete(models.RolePermission{}, "role_id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": err})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": nil, "message": "Data deleted successfully"})
}
