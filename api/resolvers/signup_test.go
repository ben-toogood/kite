package resolvers_test

import (
	"testing"

	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/usersfakes"
	"github.com/tj/assert"
)

var signupM = `
mutation {
  signup(firstName: "Alex", lastName: "B", email: "alex@m3o.com") {
    id
    firstName
  }
}
`

func TestSignup(t *testing.T) {
	u := &users.User{Id: "usr_ksjdbfks7gskduf", FirstName: "Alex"}
	testResolver.Users.(*usersfakes.FakeUsersServiceClient).CreateReturns(
		&users.CreateResponse{User: u}, nil,
	)

	res := struct {
		Signup struct {
			ID        string
			FirstName string
		}
	}{}

	test := &Test{Query: signupM, ExpectedResult: &res}
	RunQuery(t, test)
	assert.NotNil(t, res.Signup.ID)
	assert.Equal(t, u.FirstName, res.Signup.FirstName)
}
