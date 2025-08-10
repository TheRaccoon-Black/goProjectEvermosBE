package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint                `gorm:"primaryKey" json:"id"`
	UserID        uint                `gorm:"not null" json:"-"`
	AlamatKirimID uint                `gorm:"not null" json:"-"`
	HargaTotal    uint                `gorm:"not null" json:"harga_total"`
	KodeInvoice   string              `gorm:"type:varchar(50);not null;unique" json:"kode_invoice"`
	MethodBayar   string              `gorm:"type:varchar(50)" json:"method_bayar"`
	Status        string              `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	AlamatKirim   AlamatKirim         `gorm:"foreignKey:AlamatKirimID" json:"alamat_kirim"`
	DetailTrx     []TransactionDetail `gorm:"foreignKey:TransactionID" json:"detail_trx"`
	CreatedAt     time.Time           `json:"-"`
	UpdatedAt     time.Time           `json:"-"`
	DeletedAt     gorm.DeletedAt      `gorm:"index" json:"-"`
}

type TransactionDetail struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TransactionID uint      `gorm:"not null" json:"-"`
	ProductID     uint      `gorm:"not null" json:"-"`
	Kuantitas     uint      `gorm:"not null" json:"kuantitas"`
	HargaTotal    uint      `gorm:"not null" json:"harga_total"`
	Product       Product   `gorm:"foreignKey:ProductID" json:"product"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}