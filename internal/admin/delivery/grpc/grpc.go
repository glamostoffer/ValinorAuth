package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/admin/usecase"
	adminProto "github.com/glamostoffer/ValinorProtos/auth/admin_auth"
	"github.com/golang/protobuf/ptypes/empty"
)

type AdminService struct {
	uc usecase.UseCase
	adminProto.UnimplementedAdminAuthServiceServer
}

func New(uc usecase.UseCase) adminProto.AdminAuthServiceServer {
	return &AdminService{
		uc: uc,
	}
}

func (s *AdminService) AdminSignUp(
	ctx context.Context,
	request *adminProto.AdminSignUpRequest,
) (response *empty.Empty, err error) {
	return nil, nil
}
