package grpc

import (
	"context"
	"fmt"
	adminService "github.com/glamostoffer/ValinorAuth/internal/admin/delivery/grpc"
	adminUC "github.com/glamostoffer/ValinorAuth/internal/admin/usecase"
	clientService "github.com/glamostoffer/ValinorAuth/internal/client/delivery/grpc"
	clientUC "github.com/glamostoffer/ValinorAuth/internal/client/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	"github.com/glamostoffer/ValinorProtos/auth/client_auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type Server struct {
	cfg           config.Config
	grpcServer    *grpc.Server
	clientUseCase clientUC.UseCase
	adminUseCase  adminUC.UseCase
}

func NewServer(
	cfg config.Config,
	clientUseCase clientUC.UseCase,
	adminUseCase adminUC.UseCase,
) *Server {
	return &Server{
		cfg:           cfg,
		clientUseCase: clientUseCase,
		adminUseCase:  adminUseCase,
	}
}

func (s *Server) Start(_ context.Context) error {
	s.grpcServer = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: s.cfg.GRPC.Timeout,
		}),
		grpc.ChainUnaryInterceptor(),
	)
	admin_auth.RegisterAdminAuthServiceServer(s.grpcServer, adminService.New(s.adminUseCase))
	client_auth.RegisterClientAuthServiceServer(s.grpcServer, clientService.New(s.clientUseCase))

	reflection.Register(s.grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.cfg.GRPC.Host, s.cfg.GRPC.Port))
	if err != nil {
		return err
	}

	errChan := make(chan error)

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-time.After(s.cfg.StartTimeout):
		return nil
	}
}

func (s *Server) Stop(_ context.Context) error {
	stopChan := make(chan struct{})

	go func() {
		s.grpcServer.GracefulStop()
		stopChan <- struct{}{}
	}()

	select {
	case <-time.After(s.cfg.StopTimeout):
		return nil
	case <-stopChan:
		return nil
	}
}
