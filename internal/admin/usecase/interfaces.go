package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type AdminUseCase interface {
	SignUp(ctx context.Context, request model.SignUpRequest) error
	SignIn(ctx context.Context, request model.SignInRequest) (response model.SignInResponse, err error)
	GetUserDetails(ctx context.Context, clientID int64) (model.UserDetails, error)
}
