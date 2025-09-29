package models

type RolePermission struct {
	RoleId       string `gorm:"size:36;index;"`
	PermissionId string `gorm:"size:36;index;"`

	// Relations
	Role       Role       `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permission Permission `gorm:"foreignKey:PermissionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
