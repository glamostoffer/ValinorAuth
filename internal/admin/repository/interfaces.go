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
}
