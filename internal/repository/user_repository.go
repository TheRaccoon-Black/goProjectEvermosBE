package repository

import (
	"goProjectEvermos/internal/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(id uint) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *userRepository) Update(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}