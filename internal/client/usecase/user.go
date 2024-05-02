package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
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
	log := u.uc.log.With(slog.String("op", "userUseCase.SignUp"))

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", err.Error())
		return err
	}

	err = u.uc.repo.User.CreateUser(ctx, model.User{
		Username:  request.Login,
		Password:  string(passwordHashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) SignIn(
	ctx context.Context,
	request model.SignInRequest,
) (response model.SignInResponse, err error) {
	//log := u.uc.log.With(slog.String("op", "userUseCase.SignIn"))
	//
	//passwordHashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	log.Error("failed to generate password hash", err.Error())
	//	return response, err
	//}
	//
	//user, err := u.uc.repo.User.GetUserByID()

	return response, err
}

func (u *userUseCase) GetUserDetails(ctx context.Context, clientID int64) (model.UserDetails, error) {
	return model.UserDetails{}, nil
}
