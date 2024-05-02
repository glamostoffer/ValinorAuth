package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/admin/cache"
	"github.com/glamostoffer/ValinorAuth/internal/admin/repository"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
	"log/slog"
)

type UseCase struct {
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	log   *slog.Logger
	Admin AdminUseCase
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

	uc.Admin = newAdminUC(uc)

	return uc
}
