package usecase

import (
	"errors"
	"fmt"
	"goProjectEvermos/internal/domain"
	"goProjectEvermos/internal/repository"
	"goProjectEvermos/pkg/helper"
	"gorm.io/gorm"
	"github.com/gosimple/slug"
)


type RegisterUserInput struct {
	Nama      string `json:"nama"`
	KataSandi string `json:"kata_sandi"`
	NoTelp    string `json:"no_telp"`
	Email     string `json:"email"`
}

type LoginUserInput struct {
	Email     string `json:"email"`
	KataSandi string `json:"kata_sandi"`
}

type AuthUsecase interface {
	Register(input RegisterUserInput) (domain.User, error)
	Login(input LoginUserInput) (string, error)
}

type authUsecase struct {
	userRepo repository.UserRepository
	tokoRepo repository.TokoRepository
}

func NewAuthUsecase(userRepo repository.UserRepository, tokoRepo repository.TokoRepository) AuthUsecase {
	return &authUsecase{userRepo, tokoRepo}
}

func (uc *authUsecase) Register(input RegisterUserInput) (domain.User, error) {
	userByEmail, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.User{}, err 
	}
	if userByEmail.ID != 0 {
		return domain.User{}, errors.New("email sudah terdaftar")
	}

	hashedPassword, err := helper.HashPassword(input.KataSandi)
	if err != nil {
		return domain.User{}, err
	}

	newUser := domain.User{
		Nama:      input.Nama,
		Email:     input.Email,
		NoTelp:    input.NoTelp,
		KataSandi: hashedPassword,
		Role:      "user",
	}

	createdUser, err := uc.userRepo.Create(newUser)
	if err != nil {
		return domain.User{}, err
	}

	namaTokoOtomatis := slug.Make(fmt.Sprintf("toko %s", createdUser.Nama))
	newToko := domain.Toko{
		NamaToko: namaTokoOtomatis,
		UserID:   createdUser.ID,
	}
	_, err = uc.tokoRepo.Create(newToko)
	if err != nil {
		return domain.User{}, errors.New("gagal membuat toko untuk user")
	}

	return createdUser, nil
}


func (uc *authUsecase) Login(input LoginUserInput) (string, error) {
	user, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("email atau kata sandi salah")
		}
		return "", err
	}

	passwordMatch := helper.CheckPasswordHash(input.KataSandi, user.KataSandi)
	if !passwordMatch {
		return "", errors.New("email atau kata sandi salah")
	}

	token, err := helper.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}