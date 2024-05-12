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

func (r *adminRepo) GetUserByID(ctx context.Context, userID int64) (user model.User, err error) {
	log := r.adminRepo.log.With(slog.String("op", "adminRepo.GetUserByID"))

	err = r.adminRepo.db.GetContext(ctx, &user, queryGetUserByID, userID)
	if err != nil {
		log.Error("failed to get user by id", err.Error())
		return user, err
	}

	return user, nil
}

func (r *adminRepo) CreateAdmin(ctx context.Context, user model.User) error {
	log := r.adminRepo.log.With(slog.String("op", "adminRepo.CreateAdmin"))

	res, err := r.adminRepo.db.ExecContext(
		ctx,
		queryCreateAdmin,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		consts.AdminRoleID,
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

func (r *adminRepo) DeleteUser(ctx context.Context, userID int64) error {
	log := r.adminRepo.log.With(slog.String("op", "adminRepo.DeleteUser"))

	res, err := r.adminRepo.db.ExecContext(
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

func (r *adminRepo) CheckUserExists(ctx context.Context, login string) (exists bool, err error) {
	log := r.adminRepo.log.With(slog.String("op", "adminRepo.CheckUserExists"))

	err = r.adminRepo.db.GetContext(ctx, &exists, queryIsUserExists, login)
	if err != nil {
		log.Error("failed to check if user exists", err.Error())
		return exists, err
	}

	return exists, nil
}

func (r *adminRepo) GetUsers(ctx context.Context, limit, offset int64) ([]model.User, bool, error) {
	log := r.adminRepo.log.With(slog.String("op", "adminRepo.GetUsers"))

	users := make([]model.User, 0)
	err := r.adminRepo.db.SelectContext(
		ctx,
		&users,
		queryGetUsers,
		limit,
		offset,
	)
	if err != nil {
		log.Error("failed to get users from db", err.Error())
		return nil, false, err
	}

	if int64(len(users)) > limit {
		users = users[:len(users)-1]
		return users, true, nil
	} else {
		return users, false, nil
	}
}
