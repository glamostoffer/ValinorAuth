package grpc

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/client/usecase"
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
	return nil, nil
}

func (s *ClientService) SignIn(
	ctx context.Context,
	request *clientProto.SignInRequest,
) (response *clientProto.SignInResponse, err error) {
	return nil, nil
}

func (s *ClientService) GetUserDetails(
	ctx context.Context,
	request *clientProto.GetUserDetailsRequest,
) (response *clientProto.GetUserDetailsResponse, err error) {
	return nil, nil
}
