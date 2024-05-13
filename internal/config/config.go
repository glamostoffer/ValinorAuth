package config

import (
	"flag"
	rd "github.com/glamostoffer/ValinorAuth/utils/redis_connector"
	"os"
	"time"

	pg "github.com/glamostoffer/ValinorAuth/utils/pg_connector"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	StartTimeout time.Duration `yaml:"start_timeout"`
	StopTimeout  time.Duration `yaml:"stop_timeout"`
	Env          string        `yaml:"env"`
	UseCase      UseCaseConfig `yaml:"use_case"`
	GRPC         GRPCConfig    `yaml:"grpc"`
	Postgres     pg.Config     `yaml:"postgres"`
	Redis        rd.Config     `yaml:"redis"`
}

type GRPCConfig struct {
	Host    string        `yaml:"host"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type UseCaseConfig struct {
	Secret   string        `yaml:"secret"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}

func LoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag.StringVar(&configPath, "config", "", "path to config file")
		flag.Parse()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}
