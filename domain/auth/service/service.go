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

func (s *authService) Create(ctx context.Context, req *request.CreateAuthRequest) (*entity.User, error) {

	user := &entity.User{
		UniqueId:     uuid.New(),
		Name:       	req.Name,
		Email:        req.Email,
		Password:    	req.Password,
	}

	res, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *authService) Get(ctx context.Context, email string) (*entity.User, error) {

	user, err := s.repository.Get(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) GetById(ctx context.Context, id int64) (*entity.User, error) {

	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
