package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/followers"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

func TestGetFollowing(t *testing.T) {
	h := testHandler(t)

	t.Run("Invalid", func(t *testing.T) {
		rsp, err := h.GetFollowing(context.TODO(), &followers.GetFollowingRequest{})
		assert.Nil(t, rsp)

		fields := make(map[string]bool)
		for _, e := range validations.ExtractErrors(err) {
			fields[e.Field] = true
		}
		assert.True(t, fields["resource_id"])
		assert.True(t, fields["resource_type"])
	})

	t.Run("Valid", func(t *testing.T) {
		// create a test follow
		fReq := &followers.FollowRequest{
			FollowerId:    ksuid.New().String(),
			FollowerType:  followers.ResourceType_RESOURCE_TYPE_USER,
			FollowingId:   ksuid.New().String(),
			FollowingType: followers.ResourceType_RESOURCE_TYPE_USER,
		}
		_, err := h.Follow(context.TODO(), fReq)
		assert.NoError(t, err)

		// check the user was returned
		gReq := &followers.GetFollowingRequest{
			ResourceId:   fReq.FollowerId,
			ResourceType: fReq.FollowerType,
		}
		rsp, err := h.GetFollowing(context.TODO(), gReq)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.NotNil(t, rsp.Following)
		if rsp.Following == nil {
			return
		}
		f1 := rsp.Following[0]
		assert.NotNil(t, f1)
		if f1 == nil {
			return
		}
		assert.Equal(t, fReq.FollowingId, f1.Id)
		assert.Equal(t, fReq.FollowingType, f1.Type)

		// unfollow and then check again
		ufReq := &followers.UnfollowRequest{
			FollowerId:    fReq.FollowerId,
			FollowerType:  fReq.FollowerType,
			FollowingId:   fReq.FollowingId,
			FollowingType: fReq.FollowingType,
		}
		_, err = h.Unfollow(context.TODO(), ufReq)
		assert.NoError(t, err)

		gReq = &followers.GetFollowingRequest{
			ResourceId:   fReq.FollowerId,
			ResourceType: fReq.FollowerType,
		}
		rsp, err = h.GetFollowing(context.TODO(), gReq)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.Empty(t, rsp.Following)
	})
}
