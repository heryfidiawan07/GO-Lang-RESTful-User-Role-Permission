package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Role struct {
	Id string `gorm:"size:36;uniqueIndex;primaryKey"`
	Name string `gorm:"size:100;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (model *Role) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}