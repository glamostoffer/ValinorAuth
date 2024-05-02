package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type adminCache struct {
	cache *RedisCache
}

func newAdminCache(cache *RedisCache) *adminCache {
	return &adminCache{
		cache: cache,
	}
}

func (r *adminCache) SaveInviteToken(
	ctx context.Context,
	ttl time.Duration,
	token string,
) error {
	log := r.cache.log.With(slog.String("op", "adminCache.SaveInviteToken"))

	err := r.cache.rd.Set(ctx, token, true, ttl).Err()
	if err != nil {
		log.Error("failed to set token in redis", err.Error())
		return err
	}

	return nil
}

func (r *adminCache) ValidateInviteToken(
	ctx context.Context,
	token string,
) (isValid bool, err error) {
	log := r.cache.log.With(slog.String("op", "adminCache.ValidateInviteToken"))

	err = r.cache.rd.Get(ctx, token).Err()
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Error("failed to get token from redis", err.Error())
		return false, err
	}
	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	return true, err
}
