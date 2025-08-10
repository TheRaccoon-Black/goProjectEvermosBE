package domain

import (
	"time"

	"gorm.io/gorm"
)

type AlamatKirim struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UserID       uint           `gorm:"not null" json:"-"` 
	JudulAlamat  string         `gorm:"type:varchar(100);not null" json:"judul_alamat"`
	NamaPenerima string         `gorm:"type:varchar(255);not null" json:"nama_penerima"`
	NoTelp       string         `gorm:"type:varchar(20);not null" json:"no_telp"`
	DetailAlamat string         `gorm:"type:text;not null" json:"detail_alamat"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}