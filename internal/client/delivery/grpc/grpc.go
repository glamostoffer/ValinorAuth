package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/client/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"github.com/glamostoffer/ValinorAuth/utils/mapper"
	clientProto "github.com/glamostoffer/ValinorProtos/auth/client_auth"
	"github.com/golang/protobuf/ptypes/empty"
)

type ClientService struct {
	uc usecase.UseCase
	clientProto.UnimplementedClientAuthServiceServer
}

func New(uc usecase.UseCase) clientProto.ClientAuthServiceServer {
	return &ClientService{
		uc: uc,
	}
}

func (s *ClientService) SignUp(
	ctx context.Context,
	request *clientProto.SignUpRequest,
) (response *empty.Empty, err error) {
	err = s.uc.User.SignUp(
		ctx,
		model.SignUpRequest{
			Login:    request.GetLogin(),
			Password: request.GetPassword(),
		},
	)

	return nil, err
}

func (s *ClientService) SignIn(
	ctx context.Context,
	request *clientProto.SignInRequest,
) (response *clientProto.SignInResponse, err error) {
	token, err := s.uc.User.SignIn(ctx, model.SignInRequest{
		Login:    request.GetLogin(),
		Password: request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &clientProto.SignInResponse{
		Token: token,
	}, nil
}

func (s *ClientService) GetClientDetails(
	ctx context.Context,
	request *clientProto.GetClientDetailsRequest,
) (response *clientProto.GetClientDetailsResponse, err error) {
	user, err := s.uc.User.GetUserDetails(ctx, request.GetClientID())

	if err != nil {
		return nil, err
	}

	return &clientProto.GetClientDetailsResponse{
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Role:      mapper.Roles[user.Role],
	}, nil
}

func (s *ClientService) UpdateClientDetails(
	ctx context.Context,
	request *clientProto.UpdateClientDetailsRequest,
) (response *empty.Empty, err error) {
	err = s.uc.User.UpdateUserDetails(ctx, model.UpdateUserModel{
		ID:       request.GetClientID(),
		Username: request.Username,
		Password: request.Password,
	})

	return nil, err
}
