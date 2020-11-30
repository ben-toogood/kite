package handler

import (
	"context"
	"testing"

	pb "github.com/ben-toogood/kite/comments/proto"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Create(context.TODO(), &pb.CreateRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["author_id"])
		assert.True(t, fields["resource_id"])
		assert.True(t, fields["resource_type"])
		assert.True(t, fields["message"])
	})

	t.Run("Valid", func(t *testing.T) {
		req := &pb.CreateRequest{
			AuthorId:     ksuid.New().String(),
			ResourceId:   ksuid.New().String(),
			ResourceType: pb.ResourceType_Post,
			Message:      "Awesome post!",
		}
		rsp, err := h.Create(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.NotNil(t, rsp.Comment)
		if rsp.Comment == nil {
			return
		}
		assert.Contains(t, rsp.Comment.Id, "cmt_")
		assert.NotNil(t, rsp.Comment.CreatedAt)
		assert.NotNil(t, rsp.Comment.UpdatedAt)
		assert.Equal(t, req.ResourceId, rsp.Comment.ResourceId)
		assert.Equal(t, req.ResourceType, rsp.Comment.ResourceType)
		assert.Equal(t, req.Message, rsp.Comment.Message)
	})
}
