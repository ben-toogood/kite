package server

import (
	"context"

	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
)

// Create a user
func (u *Users) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, grpcError(ctx, err)
	}

	usr := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	err = model.Create(ctx, &usr)
	if err != nil {
		return nil, grpcError(ctx, err)
	}

	pbu, err := usr.Serialize()
	if err != nil {
		return nil, err
	}

	return &users.CreateResponse{User: pbu}, nil
}
