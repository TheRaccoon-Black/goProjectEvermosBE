package repository

import (
	"goProjectEvermos/internal/domain"

	"gorm.io/gorm"
)

// UserRepository mendefinisikan 'kontrak' atau fungsi apa saja yang bisa dilakukan terkait data User.
type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(id uint) (domain.User, error)
	// Tambahkan fungsi lain di sini jika dibutuhkan, misalnya Update.
}

// userRepository adalah implementasi konkret dari UserRepository.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository adalah 'constructor' untuk membuat instance userRepository baru.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create menyimpan data user baru ke database.
func (r *userRepository) Create(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

// FindByEmail mencari user berdasarkan alamat email.
func (r *userRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

// FindByID mencari user berdasarkan ID.
func (r *userRepository) FindByID(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}