package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"log/slog"
	"time"
)

type userRepo struct {
	userRepo *PgClientRepository
}

func newUserRepo(repo *PgClientRepository) *userRepo {
	return &userRepo{
		userRepo: repo,
	}
}

func (r *userRepo) GetUserByID(ctx context.Context, userID int64) (user model.User, err error) {
	log := r.userRepo.log.With(slog.String("op", "userRepo.GetUserByID"))

	err = r.userRepo.db.GetContext(ctx, &user, queryGetUserByID, userID)
	if err != nil {
		log.Error("failed to get user from table", err.Error())
		return user, err
	}

	return user, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user model.User) error {
	log := r.userRepo.log.With(slog.String("op", "userRepo.CreateUser"))

	res, err := r.userRepo.db.ExecContext(
		ctx,
		queryCreateUser,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		consts.UserRoleID,
	)
	if err != nil {
		log.Error("failed to insert user into table", err.Error())
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

func (r *userRepo) CheckUserExists(ctx context.Context, login string) (exists bool, err error) {
	log := r.userRepo.log.With(slog.String("op", "userRepo.CheckUserExists"))

	err = r.userRepo.db.GetContext(ctx, &exists, queryIsUserExists, login)
	if err != nil {
		log.Error("failed to check if user exists", err.Error())
		return exists, err
	}

	return exists, nil
}

func (r *userRepo) GetUserByLogin(ctx context.Context, login string) (user model.User, err error) {
	log := r.userRepo.log.With(slog.String("op", "userRepo.GetUserByLogin"))

	err = r.userRepo.db.GetContext(ctx, &user, queryGetUserByLogin, login)
	if err != nil {
		log.Error("failed to get user by login", err.Error())
		return user, err
	}

	return user, nil
}

func (r *userRepo) UpdateUserDetails(ctx context.Context, userInfo model.UpdateUserModel) error {
	log := r.userRepo.log.With(slog.String("op", "userRepo.UpdateUserDetails"))

	res, err := r.userRepo.db.ExecContext(
		ctx,
		queryUpdateUser,
		userInfo.ID,
		userInfo.Username,
		userInfo.Password,
		time.Now(),
	)
	if err != nil {
		log.Error("failed to update user table", err.Error())
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
