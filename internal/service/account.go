package service

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sebasvil20/templ-compl-auth/internal/config"
	"github.com/sebasvil20/templ-compl-auth/internal/model"
	"github.com/sebasvil20/templ-compl-auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IAccountService interface {
	Signup(ctx context.Context, user model.User) (*string, error)
	Login(ctx context.Context, credentials model.Credentials) (*string, error)
}

type AccountService struct {
	userRepo repository.IUserRepository
}

func NewAccountService(userRepo repository.IUserRepository) IAccountService {
	return &AccountService{
		userRepo: userRepo,
	}
}

func (s *AccountService) Signup(ctx context.Context, user model.User) (*string, error) {
	exists, err := s.userRepo.IsEmailTaken(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("email already taken")
	}

	id, err := nanoid.New(10)
	if err != nil {
		return nil, err
	}

	user.ID = id

	// Append the salt to the password
	password := []byte(user.Password + config.PasswordSalt)

	// Hash the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)
	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.generateToken(ctx, user)
}

func (s *AccountService) Login(ctx context.Context, credentials model.Credentials) (*string, error) {
	user, err := s.userRepo.GetUserBy(ctx, credentials.Email)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password+config.PasswordSalt)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return s.generateToken(ctx, *user)
}

func (s *AccountService) generateToken(ctx context.Context, user model.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	})

	tokenStr, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}
