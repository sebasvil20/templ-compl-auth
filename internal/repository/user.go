package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sebasvil20/templ-compl-auth/internal/model"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserBy(ctx context.Context, email string) (*model.User, error)
	IsEmailTaken(ctx context.Context, email string) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	_, err := r.db.
		Exec("INSERT INTO users (id, email, password, admin) VALUES (?, ?, ?, ?)", user.ID, user.Email, user.Password, user.Admin)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserBy(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	err := r.db.
		QueryRow("SELECT id, email, password, admin FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.Password, &user.Admin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	var exists bool
	row := r.db.QueryRow("SELECT EXISTS(SELECT email FROM users WHERE email = ?)", email)
	if row.Err() != nil {
		return false, row.Err()
	}
	row.Scan(&exists)
	return exists, nil
}
