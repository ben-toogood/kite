package server

import (
	"context"

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

	umap, err := model.Get(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	// serialize the result
	rsp := &users.GetResponse{
		Users: make(map[string]*users.User, len(umap)),
	}

	// Protocopy should handle maps later
	for _, u := range umap {
		var user users.User
		if err := protocopy.ToProto(u, &user); err != nil {
			return nil, err
		}
		rsp.Users[u.ID] = &user
	}

	return rsp, nil
}
