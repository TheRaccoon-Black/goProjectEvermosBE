package domain

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	NamaToko  string         `gorm:"type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string         `gorm:"type:varchar(255)" json:"url_foto"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}