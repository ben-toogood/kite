package resolvers_test

import (
	"testing"

	"github.com/ben-toogood/kite/api/resolvers"
	"github.com/ben-toogood/kite/posts"
	"github.com/ben-toogood/kite/posts/postsfakes"
)

var postsM = `
{
  getPosts() { author { id, firstName, lastName }, description, imageURL }
}
`

func TestGetPosts(t *testing.T) {
	testResolver.Posts.(*postsfakes.FakePostsServiceClient).ListReturns(
		&posts.ListResponse{},
		nil,
	)

	res := struct {
		Posts []*resolvers.Post
	}{}

	test := &Test{Query: postsM, ExpectedResult: &res}
	RunQuery(t, test)
}
