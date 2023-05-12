package repository

import (
	"context"

	"skegsTech/auth-service-go/domain/auth/entity"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
}
