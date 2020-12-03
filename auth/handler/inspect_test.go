package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/usersfakes"
	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		h := testHandler(t)
		rsp, err := h.Inspect(context.TODO(), &auth.InspectRequest{})
		assert.Nil(t, rsp)
		assert.Error(t, err)
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)
		u := &users.User{Id: "usr_ksjdbfks7gskduf", FirstName: "Alex", Email: "a@test.com"}
		h.Users.(*usersfakes.FakeUsersServiceClient).GetByEmailReturns(
			&users.GetByEmailResponse{User: u}, nil,
		)

		// generate an access token
		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{
			Email: "johndoe@gmail.com",
		})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)

		// get the jwt access token from the email
		at, err := getJWTFromTestHandlerEmail(t, h)
		assert.NoError(t, err)

		// inspect the token
		iRsp, err := h.Inspect(context.TODO(), &auth.InspectRequest{RefreshToken: at})
		assert.NoError(t, err)
		assert.NotNil(t, iRsp)
		if iRsp == nil {
			return
		}
		assert.Equal(t, u.Id, iRsp.UserId)
	})
}
