package server

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetByEmail(t *testing.T) {
	h := testServer(t)

	t.Run("MissingEmail", func(t *testing.T) {
		rsp, err := h.GetByEmail(context.TODO(), &users.GetByEmailRequest{})
		assert.Nil(t, rsp)
		assert.Error(t, err)
	})

	t.Run("RecordNotFound", func(t *testing.T) {
		rsp, err := h.GetByEmail(context.TODO(), &users.GetByEmailRequest{
			Email: "monkey@test.com",
		})
		assert.Error(t, err)
		assert.Nil(t, rsp)
	})

	t.Run("RecordFound", func(t *testing.T) {
		cReq := &users.CreateRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     uuid.New().String() + "@gmail.com",
		}
		cRsp, err := h.Create(context.TODO(), cReq)
		assert.NoError(t, err)

		gRsp, err := h.GetByEmail(context.TODO(), &users.GetByEmailRequest{
			Email: cReq.Email,
		})
		assert.NoError(t, err)
		assert.NotNil(t, gRsp)
		if gRsp == nil {
			return
		}
		assert.Equal(t, cRsp.User.Id, gRsp.User.Id)
		assert.Equal(t, cRsp.User.Email, gRsp.User.Email)
	})
}
