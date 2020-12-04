package resolvers

import (
	"github.com/ben-toogood/kite/posts"
	"github.com/graph-gophers/graphql-go"
)

type Post struct {
	p *posts.Post
}

func (p *Post) Author() *User {
	return nil
}

func (p *Post) AuthorID() graphql.ID {
	return graphql.ID(p.p.AuthorId)
}

func (p *Post) ImageURL() string {
	return p.p.ImageUrl
}

func (p *Post) Description() string {
	return p.p.Description
}
