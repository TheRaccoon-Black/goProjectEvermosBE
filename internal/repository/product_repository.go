package repository

import (
	"goProjectEvermos/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product domain.Product) (domain.Product, error)
	FindAll(filter map[string]string, page int, limit int) ([]domain.Product, error)
	FindByID(id uint) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	Delete(product domain.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product domain.Product) (domain.Product, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&product).Error; err != nil {
			return err
		}
		return nil
	})
	return product, err
}

func (r *productRepository) FindAll(filter map[string]string, page int, limit int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * limit
	query := r.db.Preload("Toko").Preload("Category").Preload("Photos").Limit(limit).Offset(offset)
	if nama := filter["nama_produk"]; nama != "" {
		query = query.Where("nama_produk LIKE ?", "%"+nama+"%")
	}
	if categoryID := filter["category_id"]; categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if tokoID := filter["toko_id"]; tokoID != "" {
		query = query.Where("toko_id = ?", tokoID)
	}
	if minHarga := filter["min_harga"]; minHarga != "" {
		query = query.Where("harga_konsumen >= ?", minHarga)
	}
	if maxHarga := filter["max_harga"]; maxHarga != "" {
		query = query.Where("harga_konsumen <= ?", maxHarga)
	}
	err := query.Find(&products).Error
	return products, err
}

func (r *productRepository) FindByID(id uint) (domain.Product, error) {
	var product domain.Product
	err := r.db.Preload("Toko").Preload("Category").Preload("Photos").First(&product, id).Error
	return product, err
}

func (r *productRepository) Update(product domain.Product) (domain.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *productRepository) Delete(product domain.Product) error {
	return r.db.Delete(&product).Error
}