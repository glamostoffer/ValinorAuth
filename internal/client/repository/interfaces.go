package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID int64) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
	CheckUserExists(ctx context.Context, login string) (exists bool, err error)
	GetUserByLogin(ctx context.Context, login string) (user model.User, err error)
	UpdateUserDetails(ctx context.Context, userInfo model.UpdateUserModel) error
}
