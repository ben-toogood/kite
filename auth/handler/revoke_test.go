package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestRevoke(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		h := testHandler(t)
		rsp, err := h.Revoke(context.TODO(), &auth.RevokeRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["user_id"])
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)

		// generate an access token
		uid := ksuid.New().String()
		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{
			UserId: uid, Email: "johndoe@gmail.com",
		})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)

		// get the jwt access token from the email
		at, err := getJWTFromTestHandlerEmail(t, h)
		assert.NoError(t, err)

		// inspect the token
		_, err = h.Inspect(context.TODO(), &auth.InspectRequest{RefreshToken: at})
		assert.NoError(t, err)

		// revoke all tokens for the user
		rRsp, err := h.Revoke(context.TODO(), &auth.RevokeRequest{UserId: uid})
		assert.NoError(t, err)
		assert.NotNil(t, rRsp)

		// inspecting the token again should fail
		_, err = h.Inspect(context.TODO(), &auth.InspectRequest{RefreshToken: at})
		assert.Error(t, err)
	})
}
