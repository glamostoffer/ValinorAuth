package usecase

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
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
