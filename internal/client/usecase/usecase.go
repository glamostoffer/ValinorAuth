package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/client/cache"
	"github.com/glamostoffer/ValinorAuth/internal/client/repository"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
	"log/slog"
)

type UseCase struct {
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	log   *slog.Logger
	User  UserUseCase
}

func New(
	cache *cache.RedisCache,
	repo *repository.PgClientRepository,
	tx *tx_manager.TxManager,
	log *slog.Logger,
) *UseCase {
	uc := &UseCase{
		cache: cache,
		repo:  repo,
		tx:    tx,
		log:   log,
	}

	uc.User = newUserUC(uc)

	return uc
}
