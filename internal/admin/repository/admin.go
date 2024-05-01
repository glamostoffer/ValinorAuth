package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type adminRepo struct {
	adminRepo *PgClientRepository
}

func newAdminRepo(repo *PgClientRepository) *adminRepo {
	return &adminRepo{
		adminRepo: repo,
	}
}

func (u *adminRepo) GetUserByID(ctx context.Context, userID int64) (model.User, error) {

	return model.User{}, nil
}

func (u *adminRepo) CreateAdmin(ctx context.Context, user model.User) error {

	return nil
}
