package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"log/slog"
)

type adminRepo struct {
	adminRepo *PgClientRepository
}

func newAdminRepo(repo *PgClientRepository) *adminRepo {
	return &adminRepo{
		adminRepo: repo,
	}
}

func (u *adminRepo) GetUserByID(ctx context.Context, userID int64) (user model.User, err error) {
	log := u.adminRepo.log.With(slog.String("op", "adminRepo.GetUserByID"))

	err = u.adminRepo.db.GetContext(ctx, &user, queryGetUserByID)
	if err != nil {
		log.Error("failed to get user by id", err.Error())
		return user, err
	}

	return user, nil
}

func (u *adminRepo) CreateAdmin(ctx context.Context, user model.User) error {
	log := u.adminRepo.log.With(slog.String("op", "adminRepo.CreateAdmin"))

	res, err := u.adminRepo.db.ExecContext(
		ctx,
		queryCreateAdmin,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		log.Error("failed to insert admin into table", err.Error())
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Error("failed to get affected rows from result", err.Error())
		return err
	}

	if rows != 1 {
		log.Error(consts.ErrInvalidAffectedRows.Error())
		return consts.ErrInvalidAffectedRows
	}

	return nil
}

func (u *adminRepo) DeleteUser(ctx context.Context, userID int64) error {
	log := u.adminRepo.log.With(slog.String("op", "adminRepo.DeleteUser"))

	res, err := u.adminRepo.db.ExecContext(
		ctx,
		queryDeleteUser,
		userID,
	)
	if err != nil {
		log.Error("failed to delete user from table", err.Error())
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Error("failed to get affected rows from result", err.Error())
		return err
	}

	if rows != 1 {
		log.Error(consts.ErrInvalidAffectedRows.Error())
		return consts.ErrInvalidAffectedRows
	}

	return nil
}
