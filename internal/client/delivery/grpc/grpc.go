package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/client/usecase"
	"github.com/glamostoffer/ValinorAuth/internal/model"
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
	return nil, nil
}

func (s *ClientService) GetClientDetails(
	ctx context.Context,
	request *clientProto.GetClientDetailsRequest,
) (response *clientProto.GetClientDetailsResponse, err error) {
	return nil, nil
}

func (s *ClientService) UpdateClientDetails(
	ctx context.Context,
	request *clientProto.UpdateClientDetailsRequest,
) (response *empty.Empty, err error) {
	return nil, nil
}
