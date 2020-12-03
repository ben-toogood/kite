package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Get(context.TODO(), &likes.GetRequest{})
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

		// get the likes for this post and another
		gRsp, err := h.Get(context.TODO(), &likes.GetRequest{
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
			ResourceIds:  []string{lReq.ResourceId, ksuid.New().String()},
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Resources)
		if gRsp.Resources == nil {
			return
		}
		assert.Len(t, gRsp.Resources, 1)

		r := gRsp.Resources[lReq.ResourceId]
		assert.NotNil(t, r)
		if r == nil {
			return
		}
		assert.Equal(t, r.Likes[0].UserId, lReq.UserId)

		// unlike the resource and check the like went away
		uReq := &likes.UnlikeRequest{
			UserId:       lReq.UserId,
			ResourceId:   lReq.ResourceId,
			ResourceType: lReq.ResourceType,
		}
		_, err = h.Unlike(context.TODO(), uReq)
		assert.NoError(t, err)

		// get the likes for this post and another
		gRsp, err = h.Get(context.TODO(), &likes.GetRequest{
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
			ResourceIds:  []string{lReq.ResourceId, ksuid.New().String()},
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.Empty(t, gRsp.Resources)
	})
}
