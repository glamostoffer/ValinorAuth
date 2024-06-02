package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type UserUseCase interface {
	ValidateToken(ctx context.Context, tokenString string) (resp model.ValidateTokenResponse, err error)
	SignUp(ctx context.Context, request model.SignUpRequest) error
	SignIn(ctx context.Context, request model.SignInRequest) (token, role string, err error)
	GetUserDetails(ctx context.Context, clientID int64) (model.User, error)
	UpdateUserDetails(ctx context.Context, userInfo model.UpdateUserModel) error
}
