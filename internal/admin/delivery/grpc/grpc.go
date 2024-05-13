package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/admin/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/utils/convert"
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

func (s *AdminService) GetListOfUsers(
	ctx context.Context,
	request *adminProto.GetListOfUsersRequest,
) (response *adminProto.GetListOfUsersResponse, err error) {
	users, hasNext, err := s.uc.Admin.GetUsers(ctx, request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, err
	}

	return &adminProto.GetListOfUsersResponse{
		Users:   convert.UsersToProto(users),
		HasNext: hasNext,
	}, nil
}

func (s *AdminService) AdminAuth(
	ctx context.Context,
	request *adminProto.AdminAuthRequest,
) (response *adminProto.AdminAuthResponse, err error) {
	out, err := s.uc.Admin.ValidateToken(ctx, request.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &adminProto.AdminAuthResponse{
		UserID: out.UserID,
		Login:  out.Login,
		Role:   out.Role,
	}, nil
}
