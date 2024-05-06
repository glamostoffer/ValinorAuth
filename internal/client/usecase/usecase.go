package usecase

import (
	"github.com/glamostoffer/ValinorAuth/internal/client/cache"
	"github.com/glamostoffer/ValinorAuth/internal/client/repository"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
	"log/slog"
	"time"
)

type Config struct {
	Secret   string        `yaml:"secret"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

type UseCase struct {
	cfg   Config
	cache *cache.RedisCache
	repo  *repository.PgClientRepository
	tx    *tx_manager.TxManager
	log   *slog.Logger
	User  UserUseCase
}

func New(
	cfg Config,
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

	uc.User = newUserUC(uc)

	return uc
}
