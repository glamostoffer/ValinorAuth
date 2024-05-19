package usecase

import (
	"context"
	"database/sql"
	"errors"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"github.com/glamostoffer/ValinorAuth/utils"
	"github.com/glamostoffer/ValinorAuth/utils/mapper"
	"github.com/golang-jwt/jwt/v5"
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
		return "", err
	}

	err = u.uc.cache.User.SaveAccessToken(ctx, token, u.uc.cfg.TokenTTL)
	if err != nil {
		return "", err
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

func (u *userUseCase) ValidateToken(ctx context.Context, tokenString string) (resp model.ValidateTokenResponse, err error) {
	log := u.uc.log.With(slog.String("op", "userUseCase.ValidateToken"))

	isValid, err := u.uc.cache.User.ValidateAccessToken(ctx, tokenString)
	if err != nil {
		return resp, err
	}
	if !isValid {
		return resp, consts.ErrInvalidAccessToken
	}

	token, err := utils.ParseJwtToken(tokenString, u.uc.cfg.Secret)
	if err != nil {
		log.Error("failed to parse token", err.Error())
		return resp, err
	}

	claims := token.Claims.(jwt.MapClaims)

	role := claims["role"].(int)
	userID := claims["id"].(int64)
	login := claims["login"].(string)

	resp.Role = mapper.Roles[role]
	resp.UserID = userID
	resp.Login = login

	return resp, nil
}
