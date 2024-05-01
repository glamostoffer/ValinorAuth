package cache

import (
	"context"
	"github.com/glamostoffer/ValinorAuth/internal/model"
	"time"
)

type UserCache interface {
	SaveSignUpRequest(
		ctx context.Context,
		token string,
		req model.SignUpRequest,
		ttl time.Duration,
	) error

	GetSignUpRequest(
		ctx context.Context,
		token string,
	) (request model.SignUpRequest, err error)
}
