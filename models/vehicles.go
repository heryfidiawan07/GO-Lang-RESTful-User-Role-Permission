package models

type Vehicles struct {
	Id int `gorm:"uniqueIndex;primaryKey"`
	PlateNumber string `gorm:"size:100;"`
	TidNumber string `gorm:"size:100;"`
	Merk string `gorm:"size:100;"`
	StnkNumber string `gorm:"size:100;"`
	HeadKir string `gorm:"size:100;"`
	Kir string `gorm:"size:100;"`
}
