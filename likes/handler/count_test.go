package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Count(context.TODO(), &likes.CountRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}

		assert.True(t, fields["resource_ids"])
		assert.True(t, fields["resource_type"])
	})

	t.Run("Valid", func(t *testing.T) {
		// like a post
		lReq := &likes.LikeRequest{
			UserId:       ksuid.New().String(),
			ResourceId:   ksuid.New().String(),
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
		}
		_, err := h.Like(context.TODO(), lReq)
		assert.NoError(t, err)

		// Count the likes for this post and another
		gRsp, err := h.Count(context.TODO(), &likes.CountRequest{
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
			ResourceIds:  []string{lReq.ResourceId, ksuid.New().String()},
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Counts)
		if gRsp.Counts == nil {
			return
		}
		assert.Equal(t, int32(1), gRsp.Counts[lReq.ResourceId])

		// unlike the resource and check the like went away
		uReq := &likes.UnlikeRequest{
			UserId:       lReq.UserId,
			ResourceId:   lReq.ResourceId,
			ResourceType: lReq.ResourceType,
		}
		_, err = h.Unlike(context.TODO(), uReq)
		assert.NoError(t, err)

		// Count the likes for this post and another
		gRsp, err = h.Count(context.TODO(), &likes.CountRequest{
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
			ResourceIds:  []string{lReq.ResourceId, ksuid.New().String()},
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Counts)
		if gRsp.Counts == nil {
			return
		}
		assert.Empty(t, gRsp.Counts)
	})
}
