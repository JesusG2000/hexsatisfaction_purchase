package grpc

import (
	"context"
)

type Checker interface {
	User(ctx context.Context, id int) (bool, error)
	Author(ctx context.Context, id int) (bool, error)
}
