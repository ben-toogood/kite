package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestUnlike(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Unlike(context.TODO(), &likes.UnlikeRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}

		assert.True(t, fields["user_id"])
		assert.True(t, fields["resource_id"])
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

		// unlike the post
		req := &likes.UnlikeRequest{
			UserId:       ksuid.New().String(),
			ResourceId:   ksuid.New().String(),
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
		}
		rsp, err := h.Unlike(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)

		// unliking the post a second time should not result in an error
		rsp, err = h.Unlike(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
	})
}
