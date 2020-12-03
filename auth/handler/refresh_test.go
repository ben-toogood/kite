package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/usersfakes"
	"github.com/stretchr/testify/assert"
)

func TestRefresh(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		h := testHandler(t)
		rsp, err := h.Refresh(context.TODO(), &auth.RefreshRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["refresh_token"])
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)
		u := &users.User{Id: "usr_ksjdbfks7gskduf", FirstName: "Alex", Email: "a@test.com"}
		h.Users.(*usersfakes.FakeUsersServiceClient).GetByEmailReturns(
			&users.GetByEmailResponse{User: u}, nil,
		)

		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{
			Email: "johndoe@gmail.com",
		})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)

		// get the jwt refresh token from the email
		at, err := getJWTFromTestHandlerEmail(t, h)
		assert.NoError(t, err)

		// refresh the token
		iRsp, err := h.Refresh(context.TODO(), &auth.RefreshRequest{RefreshToken: at})
		assert.NoError(t, err)
		assert.NotNil(t, iRsp)
		if iRsp == nil {
			return
		}
		assert.NotNil(t, iRsp.Token)
	})
}
