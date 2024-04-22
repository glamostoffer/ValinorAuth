package grpc

import (
	"github.com/glamostoffer/ValinorAuth/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg        config.Config
	grpcServer *grpc.Server
	//usecase
}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// TODO: доделать
func (s *Server) Start() error {
	s.grpcServer = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Timeout: s.cfg.GRPC.Timeout,
		}),
		grpc.ChainUnaryInterceptor(),
	)
	//admin_auth.RegisterAdminAuthServiceServer(s.grpcServer, _ /*admin_service.New()*/)
	//client_auth.RegisterClientAuthServiceServer(s.grpcServer, _ /*client_service.New()*/)

	reflection.Register(s.grpcServer)

	return nil
}
