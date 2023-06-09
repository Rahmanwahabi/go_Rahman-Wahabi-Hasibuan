package usecase

import (
	"net/http"
	"remedial/lib/database"
	"remedial/middlewares"
	"remedial/models"
	"remedial/models/payload"

	"github.com/labstack/echo/v4"
)

type UserUsecase interface {
	CreateUser(req *payload.CreateUserRequest) (*models.User, error)
	LoginUser(email, password string) (*models.User, error)
}

type userUsecase struct {
	userRepository database.UserRepository
}

func NewUserUsecase(userRepo database.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) CreateUser(req *payload.CreateUserRequest) (*models.User, error) {
	_, err := u.userRepository.GetUserByEmail(req.Email)
	if err == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Already Registered")
	}
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) LoginUser(email, password string) (*models.User, error) {
	user, err := u.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email Not Registered")
	}
	if user.Password != password {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid Password")
	}

	token, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Error Create Token")
	}
	user.Token = token
	return user, nil
}
