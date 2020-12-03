package server

import (
	"context"

	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
)

// Get users using their email
func (u *Users) GetByEmail(ctx context.Context, req *users.GetByEmailRequest) (*users.GetByEmailResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, grpcError(ctx, err)
	}

	us, err := model.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	pbu, err := us.Serialize()
	if err != nil {
		return nil, err
	}

	// serialize the result
	return &users.GetByEmailResponse{
		User: pbu,
	}, nil

}
