package main

import (
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorAuth/internal/lib/logger/pretty"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"log/slog"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info(
		"Config loaded",
		slog.Any("cfg", *cfg),
	)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case consts.EnvLocal:
		log = setupPrettySlog()
	case consts.EnvDev:
		log = setupPrettySlog()
	case consts.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := pretty.HandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
