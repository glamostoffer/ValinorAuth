package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/admin/cache"
	"github.com/glamostoffer/ValinorAuth/internal/admin/repository"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
)

type UseCase struct {
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	admin AdminUseCase
}

func New(
	cache *cache.RedisCache,
	repo *repository.PgClientRepository,
	tx *tx_manager.TxManager,
) *UseCase {
	uc := &UseCase{
		cache: cache,
		repo:  repo,
		tx:    tx,
	}

	uc.admin = newAdminUC(uc)

	return uc
}
