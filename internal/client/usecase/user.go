package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"github.com/glamostoffer/ValinorAuth/utils"
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

	exists, err := u.uc.repo.User.CheckUserExists(ctx, request.Login)
	if err != nil {
		return err
	}

	if exists {
		log.Warn("attempt to use existing login")
		return consts.ErrLoginAlreadyExists
	}

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
) (token string, err error) {
	log := u.uc.log.With(slog.String("op", "userUseCase.SignIn"))

	user, err := u.uc.repo.User.GetUserByLogin(ctx, request.Login)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return token, err
	}

	if errors.Is(sql.ErrNoRows, err) {
		log.Warn("invalid login")
		return token, consts.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Warn("invalid login")
		return token, consts.ErrInvalidCredentials
	}

	token, err = utils.NewJwtToken(user, u.uc.cfg.TokenTTL, u.uc.cfg.Secret)
	if err != nil {
		log.Error("can't create jwt token", err.Error())
		return token, err
	}

	return token, err
}

func (u *userUseCase) GetUserDetails(ctx context.Context, clientID int64) (user model.User, err error) {
	//log := u.uc.log.With(slog.String("op", "userUseCase.GetUserDetails"))

	user, err = u.uc.repo.User.GetUserByID(ctx, clientID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userUseCase) UpdateUserDetails(ctx context.Context, userInfo model.UpdateUserModel) error {
	log := u.uc.log.With(slog.String("op", "userUseCase.UpdateUserDetails"))

	if userInfo.Password != nil {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(*userInfo.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("failed to generate password hash", err.Error())
			return err
		}

		hashedPass := string(hashBytes)
		userInfo.Password = &hashedPass
	}

	err := u.uc.repo.User.UpdateUserDetails(ctx, userInfo)
	if err != nil {
		return err
	}

	return nil
}
