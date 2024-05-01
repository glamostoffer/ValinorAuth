package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/client/cache"
	"github.com/glamostoffer/ValinorAuth/internal/client/repository"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
)

type UseCase struct {
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	user  UserUseCase
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

	uc.user = newUserUC(uc)

	return uc
}
