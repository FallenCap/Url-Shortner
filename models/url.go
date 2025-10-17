package models

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	Id        uint   `gorm:"primaryKey"`
	Original  string `gorm:"type:text;not null"`
	Short     string `gorm:"size:255;uniqueIndex;not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
