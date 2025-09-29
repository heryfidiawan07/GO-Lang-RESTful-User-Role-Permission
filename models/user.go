package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id           string  `gorm:"size:36;uniqueIndex;primaryKey"`
	Name         string  `gorm:"size:100;not null"`
	Username     string  `gorm:"uniqueIndex;size:100;not null"`
	Email        string  `gorm:"uniqueIndex;size:100;not null"`
	Password     string  `gorm:"size:255;"`
	SocialId     string  `gorm:"size:100;"`
	Provider     string  `gorm:"size:100;"`
	Avatar       string  `gorm:"size:200;"`
	IsSuperadmin bool    `gorm:"default:false;"`
	RoleId       *string `gorm:"size:36;index;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// Relationship
	Role *Role `gorm:"foreignKey:RoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (model *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
