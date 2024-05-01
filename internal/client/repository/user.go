package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type userRepo struct {
	userRepo *PgClientRepository
}

func newUserRepo(repo *PgClientRepository) *userRepo {
	return &userRepo{
		userRepo: repo,
	}
}

func (u *userRepo) GetUserByID(ctx context.Context, userID int64) (model.User, error) {

	return model.User{}, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user model.User) error {

	return nil
}
