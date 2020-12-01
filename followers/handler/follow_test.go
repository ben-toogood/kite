package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestFollow(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.Follow(context.TODO(), &followers.FollowRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["follower_id"])
		assert.True(t, fields["follower_type"])
		assert.True(t, fields["following_id"])
		assert.True(t, fields["following_type"])
	})

	t.Run("Valid", func(t *testing.T) {
		req := &followers.FollowRequest{
			FollowerId:    ksuid.New().String(),
			FollowerType:  followers.ResourceType_RESOURCE_TYPE_USER,
			FollowingId:   ksuid.New().String(),
			FollowingType: followers.ResourceType_RESOURCE_TYPE_USER,
		}
		rsp, err := h.Follow(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)

		// run a second time to ensure the RPC is retry safe
		rsp, err = h.Follow(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
	})
}
