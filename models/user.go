package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type User struct {
	Id string `gorm:"size:36;uniqueIndex;primaryKey"`
	Name string `gorm:"size:100;"`
	Username string `gorm:"size:100;"`
	Email string `gorm:"unique_index;size:100;"`
	SocialId string `gorm:"size:100;"`
	Provider string `gorm:"size:100;"`
	Avatar string `gorm:"size:200;"`
	RoleId string `gorm:"size:36;index;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// relation
	// Roles  []Role `gorm:"foreignKey:id;references:role_id"`
}

func (model *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}