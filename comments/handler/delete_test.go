package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/comments"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Delete(context.TODO(), &comments.DeleteRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}

		assert.True(t, fields["id"])
	})

	t.Run("Valid", func(t *testing.T) {
		authorID := ksuid.New().String()
		resourceID := ksuid.New().String()

		// create two dummy comments
		cRsp, err := h.Create(context.TODO(), &comments.CreateRequest{
			AuthorId:     authorID,
			ResourceId:   resourceID,
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
			Message:      "Awesome post!",
		})
		assert.NoErrorf(t, err, "Error creating comment")

		// execute the query
		gRsp, err := h.Get(context.TODO(), &comments.GetRequest{
			ResourceIds:  []string{resourceID},
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
		})
		assert.NoErrorf(t, err, "Error getting comments")
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Resources)
		if gRsp.Resources == nil {
			return
		}
		r := gRsp.Resources[resourceID]
		assert.NotNil(t, r)
		if r == nil {
			return
		}
		assert.Len(t, r.Comments, 1)

		// delete the comment
		_, err = h.Delete(context.TODO(), &comments.DeleteRequest{
			Id: cRsp.Comment.Id,
		})
		assert.NoErrorf(t, err, "Error deleting comment")

		// check the comment was deleted
		gRsp, err = h.Get(context.TODO(), &comments.GetRequest{
			ResourceIds:  []string{resourceID},
			ResourceType: comments.ResourceType_RESOURCE_TYPE_POST,
		})
		assert.NoErrorf(t, err, "Error getting comments")
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Resources)
		if gRsp.Resources == nil {
			return
		}
		assert.Empty(t, gRsp.Resources[resourceID])
	})
}
