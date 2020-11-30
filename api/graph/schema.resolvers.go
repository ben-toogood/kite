package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/api/graph/generated"
	"github.com/ben-toogood/kite/api/graph/model"
	"github.com/ben-toogood/kite/users"
)

func (r *mutationResolver) Signup(ctx context.Context, input model.Signup) (*model.User, error) {
	rsp, err := r.Users.Create(ctx, &users.CreateRequest{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	})
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        rsp.User.Id,
		FirstName: rsp.User.FirstName,
		LastName:  rsp.User.LastName,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func TestSignup(t *testing.T) {

}
