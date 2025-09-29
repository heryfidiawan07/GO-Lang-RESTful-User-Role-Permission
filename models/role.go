package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Id          string `gorm:"size:36;uniqueIndex;primaryKey"`
	Name        string `gorm:"size:100;not null;uniqueIndex"`
	Desctiption string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Permissions []Permission   `gorm:"many2many:role_permissions;foreignKey:Id;joinForeignKey:RoleId;References:Id;joinReferences:PermissionId"`
}

func (model *Role) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
