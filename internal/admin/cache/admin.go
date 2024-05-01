package cache

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
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

func (r *adminCache) SaveSignUpRequest(
	ctx context.Context,
	token string,
	req model.SignUpRequest,
	ttl time.Duration,
) error {
	return nil
}

func (r *adminCache) GetSignUpRequest(
	ctx context.Context,
	token string,
) (request model.SignUpRequest, err error) {
	return model.SignUpRequest{}, err
}
