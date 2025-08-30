package repository

import (
	"goProjectEvermos/internal/domain"

	"gorm.io/gorm"
)

type TokoRepository interface {
	Create(toko domain.Toko) (domain.Toko, error)
	FindByUserID(userID uint) (domain.Toko, error)
	FindByID(tokoID uint) (domain.Toko, error)    
	Update(toko domain.Toko) (domain.Toko, error)
	FindAll(page int, limit int, namaToko string) ([]domain.Toko, error)
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

func (r *tokoRepository) FindByUserID(userID uint) (domain.Toko, error) {
	var toko domain.Toko
	err := r.db.Where("user_id = ?", userID).First(&toko).Error
	return toko, err
}

func (r *tokoRepository) FindByID(tokoID uint) (domain.Toko, error) {
	var toko domain.Toko
	err := r.db.First(&toko, tokoID).Error
	return toko, err
}

func (r *tokoRepository) Update(toko domain.Toko) (domain.Toko, error) {
	err := r.db.Save(&toko).Error
	return toko, err
}

func (r *tokoRepository) FindAll(page int, limit int, namaToko string) ([]domain.Toko, error) {
	var tokos []domain.Toko
	
	offset := (page - 1) * limit

	query := r.db.Limit(limit).Offset(offset)

	if namaToko != "" {
		query = query.Where("nama_toko LIKE ?", "%"+namaToko+"%")
	}

	err := query.Find(&tokos).Error
	return tokos, err
}