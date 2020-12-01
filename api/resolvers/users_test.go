package resolvers_test

import (
	"testing"

	"github.com/ben-toogood/kite/users"
	"github.com/ben-toogood/kite/users/usersfakes"
	"github.com/tj/assert"
)

var getM = `
query {
  user(id: "usr_ksjdbfks7gskduf") {
    id
    firstName
  }
}
`

func TestGetUser(t *testing.T) {
	u := &users.User{Id: "usr_ksjdbfks7gskduf", FirstName: "Alex"}
	testResolver.Users.(*usersfakes.FakeUsersServiceClient).GetReturns(
		&users.GetResponse{Users: map[string]*users.User{
			"usr_ksjdbfks7gskduf": u,
		}}, nil,
	)

	res := struct {
		User struct {
			ID        string
			FirstName string
		}
	}{}

	test := &Test{Query: getM, ExpectedResult: &res}
	RunQuery(t, test)
	assert.NotNil(t, res.User.ID)
	assert.Equal(t, u.FirstName, res.User.FirstName)
}
