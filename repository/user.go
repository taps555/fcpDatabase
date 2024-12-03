package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Add: Menambahkan pengguna baru ke tabel 'users'
func (u *userRepository) Add(user model.User) error {
	// Mengecek apakah username sudah ada
	var existingUser model.User
	if err := u.db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("username already exists")
	}

	// Menyimpan pengguna baru ke tabel 'users'
	if result := u.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

// CheckAvail: Memeriksa apakah pengguna dengan username dan password tertentu ada dalam tabel 'users'
func (u *userRepository) CheckAvail(user model.User) error {
	var existingUser model.User
	// Mencari pengguna berdasarkan username dan password
	if err := u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error; err != nil {
		// Jika tidak ditemukan, kembalikan error
		return fmt.Errorf("invalid username or password")
	}
	// Jika ditemukan, kembalikan nil sebagai tanda tersedia
	return nil
}
