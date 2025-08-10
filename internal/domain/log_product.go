package domain

import (
	"time"
)

type LogProduct struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	TransactionDetailID uint   `gorm:"not null" json:"-"`
	NamaProduk        string `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug              string `gorm:"type:varchar(255);not null" json:"slug"`
	HargaReseller     uint   `gorm:"not null" json:"harga_reseller"`
	HargaKonsumen     uint   `gorm:"not null" json:"harga_konsumen"`
	Deskripsi         string `gorm:"type:text" json:"deskripsi"`
	CreatedAt         time.Time `json:"-"`
}