package cache

import (
	"github.com/glamostoffer/ValinorAuth/utils/redis_connector"
)

type RedisCache struct {
	rd   *redis_connector.Connector
	user UserCache
}

func New(conn *redis_connector.Connector) *RedisCache {
	rd := &RedisCache{
		rd: conn,
	}

	rd.user = newUserCache(rd)

	return rd
}
