package domain

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	NamaCategory string         `gorm:"type:varchar(100);not null;unique" json:"nama_category"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}