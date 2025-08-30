package repository

import (
	"goProjectEvermos/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	Create(transaction domain.Transaction) (domain.Transaction, error)
	FindByUserID(userID uint) ([]domain.Transaction, error)
	FindByIDAndUserID(trxID uint, userID uint) (domain.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) Create(transaction domain.Transaction) (domain.Transaction, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		for _, detail := range transaction.DetailTrx {
			var product domain.Product
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, detail.ProductID).Error; err != nil {
				return err
			}
			if product.Stok < detail.Kuantitas {
				return gorm.ErrInvalidData
			}
			newStock := product.Stok - detail.Kuantitas
			if err := tx.Model(&product).Update("stok", newStock).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return transaction, err
}

func (r *transactionRepository) FindByUserID(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.Preload("AlamatKirim").
		Preload("DetailTrx.Product.Toko").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Photos").
		// Preload("DetailTrx.LogProduct"). 
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) FindByIDAndUserID(trxID uint, userID uint) (domain.Transaction, error) {
	var transaction domain.Transaction
	err := r.db.Preload("AlamatKirim").
		Preload("DetailTrx.Product.Toko").
		Preload("DetailTrx.Product.Category").
		Preload("DetailTrx.Product.Photos").
		// Preload("DetailTrx.LogProduct").
		Where("id = ? AND user_id = ?", trxID, userID).
		First(&transaction).Error
	return transaction, err
}