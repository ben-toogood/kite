package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/model"
	"github.com/lileio/lile/v2/protocopy"
)

// Get users using their IDs
func (u *Users) Get(ctx context.Context, req *users.GetRequest) (*users.GetResponse, error) {
	// don't query for no users
	if len(req.Ids) == 0 {
		return &users.GetResponse{}, nil
	}

	// query the database
	var usrs []model.User
	if err := u.DB.Where("id IN (?)", req.Ids).Find(&usrs).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	// serialize the result
	rsp := &users.GetResponse{
		Users: make(map[string]*users.User, len(usrs)),
	}
	for _, u := range usrs {
		var user users.User
		if err := protocopy.ToProto(u, &user); err != nil {
			return nil, err
		}
		rsp.Users[u.ID] = &user
	}
	return rsp, nil
}
