package resolvers_test

import (
	"testing"

	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/authfakes"
)

var loginM = `
mutation {
  login(email: "alex@m3o.com")
}
`

func TestLogin(t *testing.T) {
	testResolver.Auth.(*authfakes.FakeAuthServiceClient).LoginReturns(
		&auth.LoginResponse{}, nil,
	)

	res := struct {
		Login bool
	}{}

	test := &Test{Query: loginM, ExpectedResult: &res}
	RunQuery(t, test)
}
