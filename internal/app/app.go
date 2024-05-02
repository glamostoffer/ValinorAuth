package app

import (
	"context"
	"errors"
	"fmt"
	adminCache "github.com/glamostoffer/ValinorAuth/internal/admin/cache"
	adminRepository "github.com/glamostoffer/ValinorAuth/internal/admin/repository"
	adminUseCase "github.com/glamostoffer/ValinorAuth/internal/admin/usecase"
	clientCache "github.com/glamostoffer/ValinorAuth/internal/client/cache"
	clientRepository "github.com/glamostoffer/ValinorAuth/internal/client/repository"
	clientUseCase "github.com/glamostoffer/ValinorAuth/internal/client/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorAuth/internal/system/grpc"
	"github.com/glamostoffer/ValinorAuth/pkg/consts"
	"github.com/glamostoffer/ValinorAuth/utils/pg_connector"
	"github.com/glamostoffer/ValinorAuth/utils/redis_connector"
	"github.com/glamostoffer/ValinorAuth/utils/tx_manager"
	"log/slog"
)

type (
	App struct {
		cfg        config.Config
		components []component
		log        *slog.Logger
	}
	component struct {
		Service Lifecycle
		Name    string
	}
	Lifecycle interface {
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

func New(cfg config.Config, logger *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Start"))

	rd := redis_connector.New(a.cfg.Redis)
	pg := pg_connector.New(a.cfg.Postgres)

	tx := tx_manager.New(pg)

	clientRepo := clientRepository.New(pg, a.log)
	adminRepo := adminRepository.New(pg, a.log)

	clientRd := clientCache.New(rd)
	adminRd := adminCache.New(rd, a.log)

	clientUC := clientUseCase.New(
		clientRd,
		clientRepo,
		tx,
		a.log,
	)
	adminUC := adminUseCase.New(
		adminRd,
		adminRepo,
		tx,
		a.log,
	)

	grpcServer := grpc.NewServer(a.cfg, *clientUC, *adminUC)

	a.components = append(
		a.components,
		component{rd, "redis"},
		component{pg, "postgres"},
		component{tx, "tx manager"},
		component{grpcServer, "grpc server"},
	)

	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for _, c := range a.components {
			log.Info(consts.FmtStarting, slog.Any("name", c.Name))

			err = c.Service.Start(context.Background())
			if err != nil {
				log.Error(consts.FmtErrOnStarting, c.Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf("%s %s: %s", consts.FmtCannotStart, c.Name, err.Error()),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("start timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application started!")
		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	log := a.log.With(slog.String("op", "app.Stop"))
	okChan := make(chan struct{})
	errChan := make(chan error)

	go func() {
		var err error
		for i := len(a.components) - 1; i >= 0; i-- {
			log.Info(
				consts.FmtStopping,
				slog.Any("name", a.components[i].Name),
			)

			err = a.components[i].Service.Stop(ctx)
			if err != nil {
				log.Error(consts.FmtErrOnStopping, a.components[i].Name, err.Error())
				errChan <- errors.New(
					fmt.Sprintf(
						"%s %s: %s",
						consts.FmtCannotStop,
						a.components[i].Name,
						err.Error(),
					),
				)

				return
			}
		}
		okChan <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.New("stop timeout")
	case err := <-errChan:
		return err
	case <-okChan:
		log.Info("application stopped!")
		return nil
	}
}
