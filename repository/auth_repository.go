package repository

import (
	"GoSosmed/entity"

	"gorm.io/gorm"
)

// blue print untuk check email dan register, serta mengambil data user berdasarkan email
type AuthRepository interface {
	EmailExist(email string) bool
	Register(req *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
}

// mengimplementasikan Interface AuthRepository
type authRepository struct {
	db *gorm.DB
}

// contructor authRepository
func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

// impelementasi Register
func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error
	return err
}

// impelementasi checking email di database
func (r *authRepository) EmailExist(email string) bool {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return err == nil
}

// ini mencari pengguna berdasarkan email
func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return &user, err

}
