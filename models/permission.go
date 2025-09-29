package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	Id         string `gorm:"size:36;uniqueIndex;primaryKey"`
	ParentMenu string `gorm:"size:20;"`
	ParentId   string `gorm:"size:36;"`
	Name       string `gorm:"size:20;uniqueIndex;not null"`
	Alias      string `gorm:"size:20;"`
	Url        string `gorm:"size:50;"`
	Icon       string `gorm:"size:20;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (model *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
