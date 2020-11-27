package handler

import (
	"context"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users/model"
	pb "github.com/ben-toogood/kite/users/proto"
	"github.com/lileio/lile/v2/protocopy"
)

// Get users using their IDs
func (u *Users) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// don't query for no users
	if len(req.Ids) == 0 {
		return &pb.GetResponse{}, nil
	}

	// query the database
	var users []model.User
	if err := u.DB.Where("id IN (?)", req.Ids).Find(&users).Error; err != nil {
		return nil, database.TranslateErrors(err)
	}

	// serialize the result
	rsp := &pb.GetResponse{
		Users: make(map[string]*pb.User, len(users)),
	}
	for _, u := range users {
		var user pb.User
		if err := protocopy.ToProto(u, &user); err != nil {
			return nil, err
		}
		rsp.Users[u.ID] = &user
	}
	return rsp, nil
}
