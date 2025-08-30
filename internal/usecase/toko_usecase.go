package usecase

import (
	"errors"
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
)

type UpdateTokoInput struct {
	NamaToko string `form:"nama_toko"`
}

type TokoUsecase interface {
	GetMyToko(userID uint) (domain.Toko, error)
	UpdateToko(userID uint, tokoID uint, input UpdateTokoInput, fileLocation string) (domain.Toko, error)
	GetAllToko(page int, limit int, namaToko string) ([]domain.Toko, error)
	GetTokoByID(tokoID uint) (domain.Toko, error)
}

type tokoUsecase struct {
	tokoRepo repository.TokoRepository
}

func NewTokoUsecase(tokoRepo repository.TokoRepository) TokoUsecase {
	return &tokoUsecase{tokoRepo}
}

func (uc *tokoUsecase) GetMyToko(userID uint) (domain.Toko, error) {
	return uc.tokoRepo.FindByUserID(userID)
}

func (uc *tokoUsecase) UpdateToko(userID uint, tokoID uint, input UpdateTokoInput, fileLocation string) (domain.Toko, error) {
	toko, err := uc.tokoRepo.FindByID(tokoID)
	if err != nil {
		return domain.Toko{}, errors.New("toko tidak ditemukan")
	}

	if toko.UserID != userID {
		return domain.Toko{}, errors.New("anda tidak berhak mengubah toko ini")
	}

	toko.NamaToko = input.NamaToko

	if fileLocation != "" {
		toko.UrlFoto = fileLocation
	}

	updatedToko, err := uc.tokoRepo.Update(toko)
	if err != nil {
		return domain.Toko{}, err
	}

	return updatedToko, nil
}

func (uc *tokoUsecase) GetAllToko(page int, limit int, namaToko string) ([]domain.Toko, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	return uc.tokoRepo.FindAll(page, limit, namaToko)
}

func (uc *tokoUsecase) GetTokoByID(tokoID uint) (domain.Toko, error) {
	return uc.tokoRepo.FindByID(tokoID)
}