package usecase

import (
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
)


type UpdateProfileInput struct {
	Nama      string `json:"nama"`
	Tentang   string `json:"tentang"`
	Pekerjaan string `json:"pekerjaan"`
}


type UserUsecase interface {
	GetProfile(userID uint) (domain.User, error)
	UpdateProfile(userID uint, input UpdateProfileInput) (domain.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (uc *userUsecase) GetProfile(userID uint) (domain.User, error) {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (uc *userUsecase) UpdateProfile(userID uint, input UpdateProfileInput) (domain.User, error) {
	user, err := uc.userRepo.FindByID(userID)
	if err != nil {
		return domain.User{}, err
	}

	user.Nama = input.Nama
	user.Tentang = input.Tentang
	user.Pekerjaan = input.Pekerjaan

	updatedUser, err := uc.userRepo.Update(user)
	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}



