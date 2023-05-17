package service

import (
	"context"

	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/domain/auth/request"
)

type AuthService interface {
	Create(ctx context.Context, req *request.CreateAuthRequest) (*entity.User, error)
	Get(ctx context.Context, email string) (*entity.User, error)
	GetById(ctx context.Context, id int64) (*entity.User, error)
}
