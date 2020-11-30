package resolvers_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/errors"
)

var schema grqphql.Schema

func TestMain(m *testing.M) {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema = graphql.MustParseSchema(string(schemaFile), &resolvers.Resolver{}, opts...)

	os.Exit(m.Run())
}

type Test struct {
	Context        context.Context
	Query          string
	OperationName  string
	Variables      map[string]interface{}
	ExpectedResult interface{}
	ExpectedErrors []*errors.QueryError
}

func RunQuery(t *testing.T, test *Test) {
	if test.Context == nil {
		test.Context = context.Background()
	}

	result := test.Schema.Exec(test.Context, test.Query, test.OperationName, test.Variables)
	checkErrors(t, test.ExpectedErrors, result.Errors)

	// no errors so unmarshal the data into the target
	err = json.Unmarshal(result.Result, target)
	if err != nil {
		t.Logf("Response data: %s", string(result.Result))
		t.Fatalf("error unmarshaling response data: %s", err.Error())
		return t
	}
}

func checkErrors(t *testing.T, want, got []*errors.QueryError) {
	sortErrors(want)
	sortErrors(got)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("unexpected error: got %+v, want %+v", got, want)
	}
}

func sortErrors(errors []*errors.QueryError) {
	if len(errors) <= 1 {
		return
	}
	sort.Slice(errors, func(i, j int) bool {
		return fmt.Sprintf("%s", errors[i].Path) < fmt.Sprintf("%s", errors[j].Path)
	})
}
