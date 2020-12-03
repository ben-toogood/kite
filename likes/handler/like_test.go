package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/likes"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestLike(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Like(context.TODO(), &likes.LikeRequest{})
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
		req := &likes.LikeRequest{
			UserId:       ksuid.New().String(),
			ResourceId:   ksuid.New().String(),
			ResourceType: likes.ResourceType_RESOURCE_TYPE_POST,
		}

		rsp, err := h.Like(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)

		// liking the post a second time should not result in an error
		rsp, err = h.Like(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
	})
}
