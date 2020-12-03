package resolvers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"testing"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/ben-toogood/kite/auth/authfakes"
	"github.com/ben-toogood/kite/comments/commentsfakes"
	"github.com/ben-toogood/kite/users/usersfakes"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/stretchr/testify/assert"
)

var schema *graphql.Schema

var testResolver = &resolvers.Resolver{
	Users:    &usersfakes.FakeUsersServiceClient{},
	Comments: &commentsfakes.FakeCommentsServiceClient{},
	Auth:     &authfakes.FakeAuthServiceClient{},
}

type Test struct {
	Context        context.Context
	Query          string
	OperationName  string
	Variables      map[string]interface{}
	ExpectedResult interface{}
	ExpectedErrors []*errors.QueryError
}

func TestMain(m *testing.M) {
	schemaFile, err := ioutil.ReadFile("../schema.graphql")
	if err != nil {
		log.Fatal(err)
	}

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema = graphql.MustParseSchema(string(schemaFile), testResolver, opts...)

	os.Exit(m.Run())
}

func RunQuery(t *testing.T, test *Test) {
	if test.Context == nil {
		test.Context = context.Background()
	}

	result := schema.Exec(test.Context, test.Query, test.OperationName, test.Variables)
	checkErrors(t, test.ExpectedErrors, result.Errors)

	// no errors so unmarshal the data into the target
	err := json.Unmarshal(result.Data, test.ExpectedResult)
	if err != nil {
		t.Logf("Response data: %s", string(result.Data))
		t.Fatalf("error unmarshaling response data: %s", err.Error())
		return
	}
}

func checkErrors(t *testing.T, want, got []*errors.QueryError) {
	sortErrors(want)
	sortErrors(got)

	assert.EqualValues(t, got, want)
}

func sortErrors(errors []*errors.QueryError) {
	if len(errors) <= 1 {
		return
	}
	sort.Slice(errors, func(i, j int) bool {
		return fmt.Sprintf("%s", errors[i].Path) < fmt.Sprintf("%s", errors[j].Path)
	})
}
