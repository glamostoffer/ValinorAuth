package repository

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"log/slog"
)

type userRepo struct {
	userRepo *PgClientRepository
}

func newUserRepo(repo *PgClientRepository) *userRepo {
	return &userRepo{
		userRepo: repo,
	}
}

func (u *userRepo) GetUserByID(ctx context.Context, userID int64) (model.User, error) {

	return model.User{}, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user model.User) error {
	log := u.userRepo.log.With(slog.String("op", "userRepo.CreateUser"))

	res, err := u.userRepo.db.ExecContext(
		ctx,
		queryCreateUser,
		user.Username,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
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
