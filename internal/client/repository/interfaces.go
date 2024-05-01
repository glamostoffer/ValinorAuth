package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID int64) (model.User, error)
	CreateUser(ctx context.Context, user model.User) error
}
