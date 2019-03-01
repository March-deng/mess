package service

import (
	"context"
)

type store interface {
	GetUserInfo(ctx context.Context, name string) (User, error)
}
