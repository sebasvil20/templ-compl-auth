package service

import (
	"context"

	"github.com/sebasvil20/templ-compl-auth/internal/model"
	"github.com/sebasvil20/templ-compl-auth/internal/repository"
)

type IUserService interface {
	GetUserBy(ctx context.Context, email string) (*model.User, error)
}

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetUserBy(ctx context.Context, email string) (*model.User, error) {
	return s.userRepo.GetUserBy(ctx, email)
}
