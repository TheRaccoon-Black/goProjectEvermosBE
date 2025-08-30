package repository

import (
	"goProjectEvermos/internal/domain"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	Create(alamat domain.AlamatKirim) (domain.AlamatKirim, error)
	FindByUserID(userID uint) ([]domain.AlamatKirim, error)                 
	FindByIDAndUserID(alamatID uint, userID uint) (domain.AlamatKirim, error)
	Update(alamat domain.AlamatKirim) (domain.AlamatKirim, error) 
	Delete(alamat domain.AlamatKirim) error    
}

type alamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &alamatRepository{db: db}
}

func (r *alamatRepository) Create(alamat domain.AlamatKirim) (domain.AlamatKirim, error) {
	err := r.db.Create(&alamat).Error
	return alamat, err
}

func (r *alamatRepository) FindByUserID(userID uint) ([]domain.AlamatKirim, error) {
	var alamats []domain.AlamatKirim
	err := r.db.Where("user_id = ?", userID).Find(&alamats).Error
	return alamats, err
}

func (r *alamatRepository) FindByIDAndUserID(alamatID uint, userID uint) (domain.AlamatKirim, error) {
	var alamat domain.AlamatKirim
	err := r.db.Where("id = ? AND user_id = ?", alamatID, userID).First(&alamat).Error
	return alamat, err
}

func (r *alamatRepository) Update(alamat domain.AlamatKirim) (domain.AlamatKirim, error) {
	err := r.db.Save(&alamat).Error
	return alamat, err
}

func (r *alamatRepository) Delete(alamat domain.AlamatKirim) error {
	err := r.db.Delete(&alamat).Error
	return err
}
