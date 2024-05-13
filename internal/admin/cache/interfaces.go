package cache

import (
	"context"
	"time"
)

type AdminCache interface {
	SaveInviteToken(
		ctx context.Context,
		ttl time.Duration,
		token string,
	) error

	ValidateInviteToken(
		ctx context.Context,
		token string,
	) (isValid bool, err error)

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
