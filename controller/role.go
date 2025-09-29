package controller

import (
	// "fmt"
	"app/config"
	"app/models"
	"app/request"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleIndex(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": roles, "message": nil})
}

func RoleStore(c *gin.Context) {
	var valid request.RoleStore

	// Validasi input
	if err := c.ShouldBindJSON(&valid); err != nil {
		c.JSON(400, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}
	if len(valid.Permissions) == 0 {
		c.JSON(400, gin.H{"status": false, "data": nil, "message": "Permissions is required!"})
		return
	}

	// DB.Transaction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Insert role
		role := models.Role{
			Name: valid.Name,
		}
		if err := tx.Create(&role).Error; err != nil {
			return err
		}

		// Insert role_permissions
		var rolePermissions []models.RolePermission
		for _, pid := range valid.Permissions {
			rolePermissions = append(rolePermissions, models.RolePermission{
				RoleId:       role.Id,
				PermissionId: pid,
			})
		}
		if err := tx.Create(&rolePermissions).Error; err != nil {
			return err
		}

		// Setelah berhasil, preload permissions biar respons lengkap
		if err := tx.Preload("Permissions").First(&role, "id = ?", role.Id).Error; err != nil {
			return err
		}

		// Return response sukses
		c.JSON(201, gin.H{
			"status":  true,
			"message": "success",
			"data":    role,
		})
		return nil
	})

	// Kalau error transaksi
	if err != nil {
		c.JSON(500, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}
}

func RoleUpdate(c *gin.Context) {
	var valid request.RoleUpdate

	// Validasi JSON body
	if err := c.ShouldBindJSON(&valid); err != nil {
		c.JSON(400, gin.H{"status": false, "message": err.Error()})
		return
	}
	if len(valid.Permissions) == 0 {
		c.JSON(400, gin.H{"status": false, "message": "Permissions is required!"})
		return
	}

	// Cek apakah role ada
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "message": "Role not found"})
		return
	}

	// 🔍 Validasi permission_id apakah semuanya ada di tabel permissions
	var count int64
	if err := config.DB.Model(&models.Permission{}).
		Where("id IN ?", valid.Permissions).
		Count(&count).Error; err != nil {
		c.JSON(500, gin.H{"status": false, "message": "Failed to validate permissions"})
		return
	}
	if int(count) != len(valid.Permissions) {
		c.JSON(400, gin.H{"status": false, "message": "Some permissions not found"})
		return
	}

	// DB.Transaction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		// Update data role
		if err := config.DB.Model(&role).Updates(models.Role{Name: valid.Name}).Error; err != nil {
			return err
		}

		// Hapus relasi lama
		if err := config.DB.Where("role_id = ?", role.Id).Delete(&models.RolePermission{}).Error; err != nil {
			return err
		}

		// Insert relasi baru
		for _, permId := range valid.Permissions {
			rp := models.RolePermission{
				RoleId:       role.Id,
				PermissionId: permId,
			}
			if err := config.DB.Create(&rp).Error; err != nil {
				return err
			}
		}

		// Ambil data role + permissions terbaru
		if err := config.DB.Preload("Permissions").First(&role, "id = ?", role.Id).Error; err != nil {
			return err
		}

		c.JSON(200, gin.H{"status": true, "message": "success", "data": role})
		return nil
	})

	// Kalau error transaksi
	if err != nil {
		c.JSON(500, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}
}

func RoleShow(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if err := config.DB.Preload("Permissions").First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	c.JSON(200, gin.H{"status": true, "data": role, "message": nil})
}

func RoleDelete(c *gin.Context) {
	id := c.Param("id")
	var role models.Role

	if err := config.DB.First(&role, "id = ?", id).Error; err != nil {
		c.JSON(404, gin.H{"status": false, "data": nil, "message": "Data not found!"})
		return
	}

	// DB.Transaction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := config.DB.Delete(&role).Error; err != nil {
			return err
		}

		if err := config.DB.Delete(models.RolePermission{}, "role_id = ?", id).Error; err != nil {
			return err
		}

		c.JSON(200, gin.H{"status": true, "data": nil, "message": "success"})
		return nil
	})

	if err != nil {
		c.JSON(500, gin.H{"status": false, "data": nil, "message": err.Error()})
		return
	}
}
