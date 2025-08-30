package usecase

import (
	"errors"
	"fmt"
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
	"time"
)

type DetailTrxInput struct {
	ProductID uint `json:"product_id"`
	Kuantitas uint `json:"kuantitas"`
}

type CreateTransactionInput struct {
	AlamatKirimID uint             `json:"alamat_kirim_id"`
	MethodBayar   string           `json:"method_bayar"`
	DetailTrx     []DetailTrxInput `json:"detail_trx"`
}

type TransactionUsecase interface {
	CreateTransaction(userID uint, input CreateTransactionInput) (domain.Transaction, error)
	GetAllTransactions(userID uint) ([]domain.Transaction, error)
	GetTransactionByID(userID uint, trxID uint) (domain.Transaction, error)
}

type transactionUsecase struct {
	trxRepo     repository.TransactionRepository
	alamatRepo  repository.AlamatRepository
	productRepo repository.ProductRepository
}

func NewTransactionUsecase(trxRepo repository.TransactionRepository, alamatRepo repository.AlamatRepository, productRepo repository.ProductRepository) TransactionUsecase {
	return &transactionUsecase{trxRepo, alamatRepo, productRepo}
}

func (uc *transactionUsecase) CreateTransaction(userID uint, input CreateTransactionInput) (domain.Transaction, error) {
	_, err := uc.alamatRepo.FindByIDAndUserID(input.AlamatKirimID, userID)
	if err != nil {
		return domain.Transaction{}, errors.New("alamat kirim tidak valid atau bukan milik anda")
	}

	var grandTotal uint
	var detailsToSave []domain.TransactionDetail

	for _, detailInput := range input.DetailTrx {
		product, err := uc.productRepo.FindByID(detailInput.ProductID)
		if err != nil {
			return domain.Transaction{}, fmt.Errorf("produk dengan ID %d tidak ditemukan", detailInput.ProductID)
		}
		if product.Stok < detailInput.Kuantitas {
			return domain.Transaction{}, fmt.Errorf("stok produk '%s' tidak mencukupi", product.NamaProduk)
		}

		totalHargaItem := product.HargaKonsumen * detailInput.Kuantitas
		grandTotal += totalHargaItem

		detailsToSave = append(detailsToSave, domain.TransactionDetail{
			ProductID:  product.ID,
			Kuantitas:  detailInput.Kuantitas,
			HargaTotal: totalHargaItem,
		})
	}

	transaction := domain.Transaction{
		UserID:        userID,
		AlamatKirimID: input.AlamatKirimID,
		HargaTotal:    grandTotal,
		KodeInvoice:   fmt.Sprintf("INV-%d", time.Now().Unix()),
		MethodBayar:   input.MethodBayar,
		DetailTrx:     detailsToSave,
	}

	return uc.trxRepo.Create(transaction)
}

func (uc *transactionUsecase) GetAllTransactions(userID uint) ([]domain.Transaction, error) {
	return uc.trxRepo.FindByUserID(userID)
}

func (uc *transactionUsecase) GetTransactionByID(userID uint, trxID uint) (domain.Transaction, error) {
	return uc.trxRepo.FindByIDAndUserID(trxID, userID)
}