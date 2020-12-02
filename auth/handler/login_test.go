package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("Invalid", func(t *testing.T) {
		h := testHandler(t)
		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["user_id"])
		assert.True(t, fields["email"])
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)
		uid := ksuid.New().String()
		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{
			UserId: uid,
			Email:  "johndoe@gmail.com",
		})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)
		assert.Len(t, h.Sendgrid.(*sendgridMock).Messages, 1)

		// decode the JWT
		j, err := getJWTFromTestHandlerEmail(t, h)
		assert.NoError(t, err)

		// get the claims from the JWT
		data, err := base64.StdEncoding.DecodeString(strings.Split(j, ".")[1])
		assert.NoError(t, err)
		var payload jwt.MapClaims
		err = json.Unmarshal(data, &payload)
		assert.NoError(t, err)

		// check the payload is correct
		assert.Equal(t, "kite", payload["iss"])
		assert.Equal(t, uid, payload["sub"])
		return
	})
}
