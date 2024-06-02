package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type AdminRepository interface {
	GetUserByID(ctx context.Context, userID int64) (model.User, error)
	CreateAdmin(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, userID int64) error
	CheckUserExists(ctx context.Context, login string) (exists bool, err error)
	GetUsers(ctx context.Context, limit, offset int64) ([]model.User, bool, error)
	GetClientIDByLogin(ctx context.Context, login string) (int64, error)
}
