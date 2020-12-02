package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
	"github.com/lileio/lile/v2/protocopy"
	"gorm.io/gorm"
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
	rsp := users.CreateResponse{User: &users.User{}}
	if err := u.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&usr).Error; err != nil {
			return err
		}

		// serialize the result
		if err := protocopy.ToProto(usr, rsp.User); err != nil {
			return err
		}

		// publish the event
		return u.PubSub.Publish(ctx, "users.created", &users.CreatedEvent{User: rsp.User}, false)
	}); err != nil {
		return nil, database.TranslateError(err)
	}

	return &rsp, nil
}
