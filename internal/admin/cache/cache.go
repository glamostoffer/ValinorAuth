package cache

import (
	"github.com/glamostoffer/ValinorAuth/utils/redis_connector"
	"log/slog"
)

type RedisCache struct {
	rd    *redis_connector.Connector
	log   *slog.Logger
	Admin AdminCache
}

func New(conn *redis_connector.Connector, log *slog.Logger) *RedisCache {
	rd := &RedisCache{
		rd:  conn,
		log: log,
	}

	rd.Admin = newAdminCache(rd)

	return rd
}
