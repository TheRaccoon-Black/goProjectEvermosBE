package usecase

import (
	"errors"
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
	"github.com/gosimple/slug"
	"log"
)

type CreateProductInput struct {
	NamaProduk    string
	HargaReseller uint
	HargaKonsumen uint
	Stok          uint
	Deskripsi     string
	CategoryID    uint
}

type UpdateProductInput struct {
	NamaProduk    string
	HargaReseller uint
	HargaKonsumen uint
	Stok          uint
	Deskripsi     string
	CategoryID    uint
}

type ProductUsecase interface {
	CreateProduct(userID uint, input CreateProductInput, fileLocations []string) (domain.Product, error)
	GetAllProducts(filter map[string]string, page int, limit int) ([]domain.Product, error)
	GetProductByID(id uint) (domain.Product, error)
	UpdateProduct(userID uint, productID uint, input UpdateProductInput) (domain.Product, error)
	DeleteProduct(userID uint, productID uint) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
	tokoRepo    repository.TokoRepository
}

func NewProductUsecase(productRepo repository.ProductRepository, tokoRepo repository.TokoRepository) ProductUsecase {
	return &productUsecase{productRepo, tokoRepo}
}

func (uc *productUsecase) CreateProduct(userID uint, input CreateProductInput, fileLocations []string) (domain.Product, error) {
	log.Printf("MEMBUAT PRODUK UNTUK USER ID: %d", userID)
	toko, err := uc.tokoRepo.FindByUserID(userID)
	if err != nil {
		return domain.Product{}, errors.New("toko milik anda tidak ditemukan")
	}
	log.Printf("SUKSES: Ditemukan toko '%s' dengan ID: %d untuk user ID: %d", toko.NamaToko, toko.ID, userID)
	product := domain.Product{
		NamaProduk:    input.NamaProduk,
		Slug:          slug.Make(input.NamaProduk),
		HargaReseller: input.HargaReseller,
		HargaKonsumen: input.HargaKonsumen,
		Stok:          input.Stok,
		Deskripsi:     input.Deskripsi,
		TokoID:        toko.ID,
		CategoryID:    input.CategoryID,
	}

	var photos []domain.ProductPhoto
	for _, location := range fileLocations {
		photos = append(photos, domain.ProductPhoto{Url: location})
	}
	product.Photos = photos

	return uc.productRepo.Create(product)
}

func (uc *productUsecase) GetAllProducts(filter map[string]string, page int, limit int) ([]domain.Product, error) {
	if page == 0 { page = 1 }
	if limit == 0 { limit = 10 }
	return uc.productRepo.FindAll(filter, page, limit)
}

func (uc *productUsecase) GetProductByID(id uint) (domain.Product, error) {
	return uc.productRepo.FindByID(id)
}

func (uc *productUsecase) UpdateProduct(userID uint, productID uint, input UpdateProductInput) (domain.Product, error) {
	// Verifikasi kepemilikan produk
	toko, err := uc.tokoRepo.FindByUserID(userID)
	if err != nil {
		return domain.Product{}, errors.New("toko milik anda tidak ditemukan")
	}
	product, err := uc.productRepo.FindByID(productID)
	if err != nil {
		return domain.Product{}, errors.New("produk tidak ditemukan")
	}
	if product.TokoID != toko.ID {
		return domain.Product{}, errors.New("anda tidak berhak mengubah produk ini")
	}

	product.NamaProduk = input.NamaProduk
	product.Slug = slug.Make(input.NamaProduk)
	product.HargaReseller = input.HargaReseller
	product.HargaKonsumen = input.HargaKonsumen
	product.Stok = input.Stok
	product.Deskripsi = input.Deskripsi
	product.CategoryID = input.CategoryID

	return uc.productRepo.Update(product)
}

func (uc *productUsecase) DeleteProduct(userID uint, productID uint) error {
	// Verifikasi kepemilikan produk
	toko, err := uc.tokoRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("toko milik anda tidak ditemukan")
	}
	product, err := uc.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}
	if product.TokoID != toko.ID {
		return errors.New("anda tidak berhak menghapus produk ini")
	}

	return uc.productRepo.Delete(product)
}