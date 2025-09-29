package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
	Id        string `gorm:"size:36;uniqueIndex;primaryKey"`
	Revoked   bool   `gorm:"size:1;"`
	ExpiredAt time.Time
	UserId    string `gorm:"size:36;index;not null;"`

	// Relationship
	User User `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (model *RefreshToken) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	model.Id = uuid.NewString()
	return
}
