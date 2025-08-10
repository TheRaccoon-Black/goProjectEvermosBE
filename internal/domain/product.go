package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	NamaProduk    string         `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug          string         `gorm:"type:varchar(255);not null;unique" json:"slug"`
	HargaReseller uint           `gorm:"not null" json:"harga_reseller"`
	HargaKonsumen uint           `gorm:"not null" json:"harga_konsumen"`
	Stok          uint           `gorm:"not null;default:0" json:"stok"`
	Deskripsi     string         `gorm:"type:text" json:"deskripsi"`
	TokoID        uint           `gorm:"not null" json:"-"`
	CategoryID    uint           `gorm:"not null" json:"-"`
	Toko          Toko           `gorm:"foreignKey:TokoID" json:"toko"`
	Category      Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Photos        []ProductPhoto `gorm:"foreignKey:ProductID" json:"photos"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}