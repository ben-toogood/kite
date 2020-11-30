package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Get(context.TODO(), &comments.GetRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}

		assert.True(t, fields["resource_ids"])
		assert.True(t, fields["resource_type"])
	})

	t.Run("Valid", func(t *testing.T) {
		authorID := ksuid.New().String()
		resourceIDOne := ksuid.New().String()
		resourceIDTwo := ksuid.New().String()

		// create two dummy comments
		_, err := h.Create(context.TODO(), &comments.CreateRequest{
			AuthorId:     authorID,
			ResourceId:   resourceIDOne,
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
			Message:      "Awesome post!",
		})
		assert.NoErrorf(t, err, "Error creating comment")

		_, err = h.Create(context.TODO(), &comments.CreateRequest{
			AuthorId:     authorID,
			ResourceId:   resourceIDTwo,
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
			Message:      "Awesome post!",
		})
		assert.NoErrorf(t, err, "Error creating comment")

		// execute the query
		rsp, err := h.Get(context.TODO(), &comments.GetRequest{
			ResourceIds:  []string{resourceIDOne, resourceIDTwo},
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
		})
		assert.NoErrorf(t, err, "Error getting comments")
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.NotNil(t, rsp.Resources)
		if rsp.Resources == nil {
			return
		}

		r1 := rsp.Resources[resourceIDOne]
		r2 := rsp.Resources[resourceIDTwo]
		assert.NotNil(t, r1)
		assert.NotNil(t, r2)
		if r1 == nil || r2 == nil {
			return
		}
		assert.Len(t, r1.Comments, 1)
		assert.Len(t, r2.Comments, 1)

		// check all the comment info was returned
		c := r1.Comments[0]
		if c == nil {
			return
		}
		assert.NotEmpty(t, c.Id)
		assert.NotZero(t, c.CreatedAt)
		assert.Equal(t, "Awesome post!", c.Message)
		assert.Equal(t, resourceIDOne, c.ResourceId)
		assert.Equal(t, comments.ResourceType_RESOURCE_TYPE_POST, c.ResourceType)
		assert.Equal(t, authorID, c.AuthorId)
	})
}
