package repository

import (
	"goProjectEvermos/internal/domain"

	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(toko domain.Toko) (domain.Toko, error)
}

type tokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db: db}
}

func (r *tokoRepository) Create(toko domain.Toko) (domain.Toko, error) {
	err := r.db.Create(&toko).Error
	return toko, err
}