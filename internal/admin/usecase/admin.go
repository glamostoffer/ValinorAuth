package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"github.com/glamostoffer/ValinorAuth/utils"
	"github.com/glamostoffer/ValinorAuth/utils/mapper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type adminUseCase struct {
	uc *UseCase
}

func newAdminUC(useCase *UseCase) *adminUseCase {
	return &adminUseCase{
		uc: useCase,
	}
}

func (u *adminUseCase) SignUp(ctx context.Context, request model.AdminSignUpRequest) error {
	log := u.uc.log.With(slog.String("op", "adminUseCase.SignUp"))

	exists, err := u.uc.repo.Admin.CheckUserExists(ctx, request.Login)
	if err != nil {
		return err
	}

	if exists {
		log.Warn("attempt to use existing login")
		return consts.ErrLoginAlreadyExists
	}

	isValid, err := u.uc.cache.Admin.ValidateInviteToken(ctx, request.InviteToken)
	if err != nil {
		return err
	}
	if !isValid {
		log.Info(consts.ErrInvalidInviteToken.Error())
		return consts.ErrInvalidInviteToken
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", err.Error())
		return err
	}

	err = u.uc.repo.Admin.CreateAdmin(ctx, model.User{
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

func (u *adminUseCase) BanUser(ctx context.Context, clientID int64) error {
	return u.uc.repo.Admin.DeleteUser(ctx, clientID)
}

func (u *adminUseCase) CreateInviteToken(ctx context.Context, ttl int64) (string, error) {
	//log := u.uc.log.With(slog.String("op", "adminUseCase.CreateInviteToken"))

	token := uuid.NewString()

	ttlDuration := time.Unix(ttl, 0).Sub(time.Now())

	err := u.uc.cache.Admin.SaveInviteToken(ctx, ttlDuration, token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *adminUseCase) GetUsers(ctx context.Context, limit, offset int64) ([]model.User, bool, error) {
	//log := u.uc.log.With(slog.String("op", "adminUseCase.GetUsers"))

	return u.uc.repo.Admin.GetUsers(ctx, limit, offset)
}

func (u *adminUseCase) ValidateToken(ctx context.Context, tokenString string) (resp model.ValidateTokenResponse, err error) {
	log := u.uc.log.With(slog.String("op", "adminUseCase.ValidateToken"))

	isValid, err := u.uc.cache.Admin.ValidateAccessToken(ctx, tokenString)
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

	role := int(claims["role"].(float64))
	if role != consts.AdminRoleID {
		return resp, consts.ErrInvalidAccessToken
	}

	userID := int64(claims["id"].(float64))
	login := claims["login"].(string)

	resp.Role = mapper.Roles[role]
	resp.UserID = userID
	resp.Login = login

	return resp, nil
}

func (u *adminUseCase) GetClientIDByLogin(ctx context.Context, login string) (int64, error) {
	return u.uc.repo.Admin.GetClientIDByLogin(ctx, login)
}
