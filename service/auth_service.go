package service

import (
	"GoSosmed/dto"
	"GoSosmed/entity"
	"GoSosmed/errorhandler"
	"GoSosmed/helper"
	"GoSosmed/repository"
)

// mendefinisikan kontrak untuk layanan autentikasi
type AuthService interface {
	Register(req *dto.ResgisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

// implementasikan interface AuthService
type authService struct {
	repository repository.AuthRepository
}

// constructor instance authService
func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

// check email sudah terdaftar dalam database atau tidak
func (s *authService) Register(req *dto.ResgisterRequest) error {
	if EmailExist := s.repository.EmailExist(req.Email); EmailExist {
		return &errorhandler.BadRequestError{
			Message: "email alredy registered",
		}
	}
	// check kecocokan password
	if req.Password != req.PasswordConfirmation {
		return &errorhandler.BadRequestError{
			Message: "password not match",
		}
	}
	//Menghasilkan hash dari password yang di kirim user
	passwordHahs, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}
	//menyimpan data dari request pengguna yang telah di hash
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: passwordHahs,
		Gender:   req.Gender,
	}
	//Menyimpan entitas User ke dalam database
	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

// Login yang digunakan untuk autentikasi pengguna
func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse
	//Mencari pengguna berdasarkan email
	user, err := s.repository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "wrong email or password",
		}
	}
	//Memverifikasi apakah password cocok
	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: "wrong email or password ",
		}
	}
	//Menghasilkan token JWT untuk user
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &data, nil
}
