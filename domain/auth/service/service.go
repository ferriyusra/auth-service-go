package service

import (
	"context"

	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/domain/auth/repository"
	"skegsTech/auth-service-go/domain/auth/request"

	"github.com/google/uuid"
)

type authService struct {
	repository repository.UserRepository
}

func NewService(repo repository.UserRepository) AuthService {
	return &authService{
		repository: repo,
	}
}

func (s *authService) Register(ctx context.Context, req *request.CreateAuthRequest) (*entity.User, error) {

	user := &entity.User{
		UniqueId:     uuid.New(),
		Name:       	req.Name,
		Email:        req.Email,
		Password:    	req.Password,
	}

	res, err := s.repository.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
