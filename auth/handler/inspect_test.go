package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/auth"
	"github.com/segmentio/ksuid"
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
		iRsp, err := h.Inspect(context.TODO(), &auth.InspectRequest{RefreshToken: at})
		assert.NoError(t, err)
		assert.NotNil(t, iRsp)
		if iRsp == nil {
			return
		}
		assert.Equal(t, uid, iRsp.UserId)
	})
}
