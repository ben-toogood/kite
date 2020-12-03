package resolvers_test

import (
	"testing"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/ben-toogood/kite/auth"
	"github.com/ben-toogood/kite/auth/authfakes"
	"github.com/stretchr/testify/assert"
)

var refreshM = `
mutation {
  tokens:refreshTokens(refreshToken: "eyal8s7dhaol8s97dbla98s7dbla98") {accessToken, refreshToken}
}
`

func TestRefreshTokens(t *testing.T) {
	testResolver.Auth.(*authfakes.FakeAuthServiceClient).RefreshReturns(
		&auth.RefreshResponse{Token: &auth.Token{AccessToken: "at", RefreshToken: "rt"}}, nil,
	)

	res := struct {
		Tokens resolvers.Tokens
	}{}

	test := &Test{Query: refreshM, ExpectedResult: &res}
	RunQuery(t, test)
	assert.NotNil(t, res.Tokens.AccessToken)
	assert.Equal(t, res.Tokens.AccessToken, "at")
	assert.NotNil(t, res.Tokens.RefreshToken)
	assert.Equal(t, res.Tokens.RefreshToken, "rt")
}
