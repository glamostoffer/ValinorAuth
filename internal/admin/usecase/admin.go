package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
)

type adminUseCase struct {
	uc *UseCase
}

func newAdminUC(useCase *UseCase) *adminUseCase {
	return &adminUseCase{
		uc: useCase,
	}
}

func (u *adminUseCase) SignUp(ctx context.Context, request model.SignUpRequest) error {
	return nil
}

func (u *adminUseCase) SignIn(
	ctx context.Context,
	request model.SignInRequest,
) (response model.SignInResponse, err error) {
	return model.SignInResponse{}, err
}

func (u *adminUseCase) GetUserDetails(ctx context.Context, clientID int64) (model.UserDetails, error) {
	return model.UserDetails{}, nil
}
