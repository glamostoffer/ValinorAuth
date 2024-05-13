package cache

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"time"
)

type UserCache interface {
	SaveInviteToken(
		ctx context.Context,
		token string,
		req model.SignUpRequest,
		ttl time.Duration,
	) error

	GetSignUpRequest(
		ctx context.Context,
		token string,
	) (request model.SignUpRequest, err error)

	SaveAccessToken(
		ctx context.Context,
		token string,
		ttl time.Duration,
	) error

	ValidateAccessToken(
		ctx context.Context,
		token string,
	) (bool, error)
}
