package usecase

import (
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
)

type CreateAlamatInput struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}


type UpdateAlamatInput struct {
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}


type AlamatUsecase interface {
	CreateAlamat(userID uint, input CreateAlamatInput) (domain.AlamatKirim, error)
	GetAllAlamat(userID uint) ([]domain.AlamatKirim, error)                 	
	GetAlamatByID(userID uint, alamatID uint) (domain.AlamatKirim, error)
	UpdateAlamat(userID uint, alamatID uint, input UpdateAlamatInput) (domain.AlamatKirim, error) 
	DeleteAlamat(userID uint, alamatID uint) error   
}

type alamatUsecase struct {
	alamatRepo repository.AlamatRepository
}

func NewAlamatUsecase(alamatRepo repository.AlamatRepository) AlamatUsecase {
	return &alamatUsecase{alamatRepo}
}

func (uc *alamatUsecase) CreateAlamat(userID uint, input CreateAlamatInput) (domain.AlamatKirim, error) {
	alamat := domain.AlamatKirim{
		JudulAlamat:  input.JudulAlamat,
		NamaPenerima: input.NamaPenerima,
		NoTelp:       input.NoTelp,
		DetailAlamat: input.DetailAlamat,
		UserID:       userID, 
	}

	newAlamat, err := uc.alamatRepo.Create(alamat)
	if err != nil {
		return domain.AlamatKirim{}, err
	}

	return newAlamat, nil
}

func (uc *alamatUsecase) GetAllAlamat(userID uint) ([]domain.AlamatKirim, error) {
	return uc.alamatRepo.FindByUserID(userID)
}

func (uc *alamatUsecase) GetAlamatByID(userID uint, alamatID uint) (domain.AlamatKirim, error) {
	return uc.alamatRepo.FindByIDAndUserID(alamatID, userID)
}

// Tambahkan fungsi baru UpdateAlamat
func (uc *alamatUsecase) UpdateAlamat(userID uint, alamatID uint, input UpdateAlamatInput) (domain.AlamatKirim, error) {
	// 1. Verifikasi kepemilikan alamat terlebih dahulu
	alamat, err := uc.alamatRepo.FindByIDAndUserID(alamatID, userID)
	if err != nil {
		return domain.AlamatKirim{}, err 
	}

	alamat.JudulAlamat = input.JudulAlamat
	alamat.NamaPenerima = input.NamaPenerima
	alamat.NoTelp = input.NoTelp
	alamat.DetailAlamat = input.DetailAlamat

	updatedAlamat, err := uc.alamatRepo.Update(alamat)
	if err != nil {
		return domain.AlamatKirim{}, err
	}

	return updatedAlamat, nil
}

func (uc *alamatUsecase) DeleteAlamat(userID uint, alamatID uint) error {
	alamat, err := uc.alamatRepo.FindByIDAndUserID(alamatID, userID)
	if err != nil {
		return err
	}

	return uc.alamatRepo.Delete(alamat)
}