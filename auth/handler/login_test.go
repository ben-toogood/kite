package handler

import (
	"context"
	"testing"

	"github.com/dgrijalva/jwt-go"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/usersfakes"
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
		assert.True(t, fields["email"])
	})

	t.Run("Valid", func(t *testing.T) {
		h := testHandler(t)

		u := &users.User{Id: "usr_ksjdbfks7gskduf", FirstName: "Alex", Email: "a@test.com"}
		h.Users.(*usersfakes.FakeUsersServiceClient).GetByEmailReturns(
			&users.GetByEmailResponse{User: u}, nil,
		)

		rsp, err := h.Login(context.TODO(), &auth.LoginRequest{
			Email: u.Email,
		})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)
		assert.Len(t, h.Sendgrid.(*sendgridMock).Messages, 1)

		// decode the JWT
		j, err := getJWTFromTestHandlerEmail(t, h)
		assert.NoError(t, err)

		jwt.Parse(j, func(tk *jwt.Token) (interface{}, error) {
			assert.Equal(t, "kite", tk.Claims.(jwt.MapClaims)["iss"])
			assert.Equal(t, u.Id, tk.Claims.(jwt.MapClaims)["sub"])
			return nil, nil
		})
		return
	})
}
