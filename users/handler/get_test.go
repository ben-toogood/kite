package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingIDs", func(t *testing.T) {
		rsp, err := h.Get(context.TODO(), &users.GetRequest{})
		assert.NotNil(t, rsp)
		assert.NoError(t, err)
	})

	t.Run("RecordNotFound", func(t *testing.T) {
		rsp, err := h.Get(context.TODO(), &users.GetRequest{
			Ids: []string{"usr-one"},
		})
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.NotNil(t, rsp.Users)
		if rsp.Users == nil {
			return
		}
		assert.Empty(t, rsp.Users)
	})

	t.Run("RecordFound", func(t *testing.T) {
		cReq := &users.CreateRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     uuid.New().String() + "@gmail.com",
		}
		cRsp, err := h.Create(context.TODO(), cReq)
		assert.NoError(t, err)

		gRsp, err := h.Get(context.TODO(), &users.GetRequest{
			Ids: []string{cRsp.User.Id},
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.NotNil(t, gRsp.Users)
		if gRsp.Users == nil {
			return
		}
		assert.Len(t, gRsp.Users, 1)
		u := gRsp.Users[cRsp.User.Id]
		assert.NotNil(t, u)
		if u == nil {
			return
		}
		assert.Equal(t, cRsp.User.Id, u.Id)
		assert.Equal(t, cReq.FirstName, u.FirstName)
		assert.Equal(t, cReq.LastName, u.LastName)
		assert.Equal(t, cReq.Email, u.Email)
	})
}
