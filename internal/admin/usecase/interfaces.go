package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type AdminUseCase interface {
	ValidateToken(ctx context.Context, token string) (model.ValidateTokenResponse, error)
	SignUp(ctx context.Context, request model.AdminSignUpRequest) error
	BanUser(ctx context.Context, clientID int64) error
	CreateInviteToken(ctx context.Context, ttl int64) (string, error)
	GetUsers(ctx context.Context, limit, offset int64) ([]model.User, bool, error)
}
