package main

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/app"
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"github.com/glamostoffer/ValinorAuth/utils/logger/pretty"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.LoadConfig()

	log := setupLogger(cfg.Env)

	log.Info(
		"Config loaded",
		slog.Any("cfg", *cfg),
	)

	application := app.New(*cfg, log)

	go func() {
		if err := application.Start(context.Background()); err != nil {
			panic(err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	stopCtx, stopCancel := context.WithTimeout(context.Background(), cfg.StopTimeout)
	defer stopCancel()

	if err := application.Stop(stopCtx); err != nil {
		panic(err.Error())
	}
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
