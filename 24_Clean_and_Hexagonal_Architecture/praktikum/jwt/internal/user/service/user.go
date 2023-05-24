package service

import (
	"code/internal/user/dto"
	"context"
)

type UserService interface {
	FindAll(ctx context.Context) (dto.UsersResponse, error)
	CreateUser(user dto.UserRequest, ctx context.Context) error
	Login(user dto.UserRequest, ctx context.Context) (string, error)
}
