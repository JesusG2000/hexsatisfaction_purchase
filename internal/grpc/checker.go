package grpc

import (
	"context"

	"github.com/JesusG2000/hexsatisfaction/pkg/grpc/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type CheckerService struct {
	client api.ExistanceClient
}

func NewCheckerService(addr string) (*CheckerService, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "couldn't dial connection with gprc")
	}

	return &CheckerService{api.NewExistanceClient(conn)}, nil
}

func (c CheckerService) User(ctx context.Context, id int) (bool, error) {
	res, err := c.client.User(ctx, &api.IsUserExistRequest{Id: int32(id)})
	if err != nil {
		return false, errors.Wrap(err, "couldn't check user existence")
	}

	return res.Exist, nil
}
func (c CheckerService) Author(ctx context.Context, id int) (bool, error) {
	res, err := c.client.Author(ctx, &api.IsAuthorExistRequest{Id: int32(id)})
	if err != nil {
		return false, errors.Wrap(err, "couldn't check author existence")
	}

	return res.Exist, nil
}
