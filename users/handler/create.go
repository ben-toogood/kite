package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
	"github.com/lileio/lile/v2/protocopy"
)

// Create a user
func (u *Users) Create(ctx context.Context, req *users.CreateRequest) (*users.CreateResponse, error) {
	// validate the request
	if err := req.Validate(); err != nil {
		return nil, validations.NewError(ctx, err)
	}

	// construct the object and write it to the database
	usr := model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}
	if err := u.DB.Create(&usr).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	// serialize the result
	rsp := users.CreateResponse{User: &users.User{}}
	if err := protocopy.ToProto(usr, rsp.User); err != nil {
		return nil, err
	}
	return &rsp, nil
}
