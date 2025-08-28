package domain

import (
	"time"
	
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nama         string    `gorm:"type:varchar(255);not null" json:"nama"`
	KataSandi    string    `gorm:"type:varchar(255);not null" json:"-"` 
	NoTelp       string    `gorm:"type:varchar(20);not null;unique" json:"no_telp"`
	TanggalLahir string    `gorm:"type:varchar(50)" json:"tanggal_lahir,omitempty"`
	Tentang      string    `gorm:"type:text" json:"tentang,omitempty"`
	Pekerjaan    string    `gorm:"type:varchar(100)" json:"pekerjaan,omitempty"`
	Email        string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	IDProvinsi   string    `gorm:"type:varchar(10)" json:"id_provinsi,omitempty"`
	IDKota       string    `gorm:"type:varchar(10)" json:"id_kota,omitempty"`
	Role         string    `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"` 
}