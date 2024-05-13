package cache

import (
	"context"
	"errors"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type userCache struct {
	cache *RedisCache
}

func newUserCache(cache *RedisCache) *userCache {
	return &userCache{
		cache: cache,
	}
}

func (r *userCache) SaveInviteToken(
	ctx context.Context,
	token string,
	req model.SignUpRequest,
	ttl time.Duration,
) error {
	return nil
}

func (r *userCache) GetSignUpRequest(
	ctx context.Context,
	token string,
) (request model.SignUpRequest, err error) {
	return model.SignUpRequest{}, err
}

func (r *userCache) SaveAccessToken(
	ctx context.Context,
	token string,
	ttl time.Duration,
) error {
	log := r.cache.log.With(slog.String("op", "userCache.SaveAccessToken"))

	err := r.cache.rd.Set(ctx, token, true, ttl).Err()
	if err != nil {
		log.Error("failed to set token in redis", err.Error())
		return err
	}
	return nil
}

func (r *userCache) ValidateAccessToken(
	ctx context.Context,
	token string,
) (isValid bool, err error) {
	log := r.cache.log.With(slog.String("op", "userCache.GetAccessToken"))

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
