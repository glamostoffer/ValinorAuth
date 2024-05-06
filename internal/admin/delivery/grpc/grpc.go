package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/admin/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/model"
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
	err = s.uc.Admin.SignUp(
		ctx,
		model.AdminSignUpRequest{
			Login:       request.GetLogin(),
			Password:    request.GetPassword(),
			InviteToken: request.GetInviteToken(),
		},
	)

	return nil, err
}

func (s *AdminService) BanUser(
	ctx context.Context,
	request *adminProto.BanUserRequest,
) (response *empty.Empty, err error) {
	err = s.uc.Admin.BanUser(ctx, request.GetUserID())

	return nil, err
}

func (s *AdminService) CreateInviteToken(
	ctx context.Context,
	request *adminProto.CreateInviteTokenRequest,
) (response *adminProto.CreateInviteTokenResponse, err error) {
	token, err := s.uc.Admin.CreateInviteToken(ctx, request.GetTtl())
	if err != nil {
		return nil, err
	}

	return &adminProto.CreateInviteTokenResponse{
		Token: token,
	}, nil
}
