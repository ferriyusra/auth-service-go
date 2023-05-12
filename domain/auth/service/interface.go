package service

import (
	"context"

	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/domain/auth/request"
)

type AuthService interface {
	Register(ctx context.Context, req *request.CreateAuthRequest) (*entity.User, error)
}
