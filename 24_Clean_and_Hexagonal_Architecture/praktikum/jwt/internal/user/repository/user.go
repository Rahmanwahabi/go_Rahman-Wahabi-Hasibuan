package repository

import (
	"code/pkg/entity"
	"context"
)

type UserRepository interface {
	FindAll(ctx context.Context) (entity.Users, error)
	CreateUser(user *entity.User, ctx context.Context) error
	FindByEmail(email string, ctx context.Context) (*entity.User, error)
}
