package domain

import (
	"time"

	"gorm.io/gorm"
)

type ProductPhoto struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Url       string         `gorm:"type:varchar(255);not null" json:"url"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}