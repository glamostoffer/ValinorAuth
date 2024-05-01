package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type userUseCase struct {
	uc *UseCase
}

func newUserUC(useCase *UseCase) *userUseCase {
	return &userUseCase{
		uc: useCase,
	}
}

func (u *userUseCase) SignUp(ctx context.Context, request model.SignUpRequest) error {
	return nil
}

func (u *userUseCase) SignIn(
	ctx context.Context,
	request model.SignInRequest,
) (response model.SignInResponse, err error) {
	return model.SignInResponse{}, err
}

func (u *userUseCase) GetUserDetails(ctx context.Context, clientID int64) (model.UserDetails, error) {
	return model.UserDetails{}, nil
}
