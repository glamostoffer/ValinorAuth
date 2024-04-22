package main

import (
	"fmt"
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info(fmt.Sprintf("cfg: %+v", *cfg))
	log.Info("Config loaded")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case consts.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case consts.EnvDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case consts.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
