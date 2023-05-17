package repository

import (
	"context"

	"skegsTech/auth-service-go/domain/auth/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Get(ctx context.Context, email string) (*entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
}
