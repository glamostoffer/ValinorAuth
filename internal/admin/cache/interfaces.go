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
}
