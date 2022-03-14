package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type RefreshToken struct {
	Id string `gorm:"size:36;uniqueIndex;primaryKey"`
	Revoked bool `gorm:"size:1;"`
	ExpiredAt time.Time
	UserId string `gorm:"size:36;index;"`
}

func (model *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
