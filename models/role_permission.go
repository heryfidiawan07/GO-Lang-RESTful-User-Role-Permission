package models

type RolePermission struct {
	RoleId string `gorm:"size:36;index;"`
	PermissionId string `gorm:"size:36;index;"`
}
