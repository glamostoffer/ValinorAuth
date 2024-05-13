package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/admin/cache"
	"github.com/glamostoffer/ValinorAuth/internal/admin/repository"
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
	"log/slog"
)

type UseCase struct {
	cfg   config.UseCaseConfig
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	log   *slog.Logger
	Admin AdminUseCase
}

func New(
	cfg config.UseCaseConfig,
	cache *cache.RedisCache,
	repo *repository.PgClientRepository,
	tx *tx_manager.TxManager,
	log *slog.Logger,
) *UseCase {
	uc := &UseCase{
		cfg:   cfg,
		cache: cache,
		repo:  repo,
		tx:    tx,
		log:   log,
	}

	uc.Admin = newAdminUC(uc)

	return uc
}
