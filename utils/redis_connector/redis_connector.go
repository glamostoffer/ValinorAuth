package redis_connector

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Connector struct {
	*redis.Client
	cfg Config
}

func New(cfg Config) *Connector {
	return &Connector{
		cfg: cfg,
	}
}

func (c *Connector) Start(_ context.Context) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.cfg.Host, c.cfg.Port),
		Password: c.cfg.Password,
		DB:       c.cfg.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return err
	}

	c.Client = rdb

	return nil
}

func (c *Connector) Stop(_ context.Context) error {
	return c.Close()
}
