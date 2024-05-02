package cache

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
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
