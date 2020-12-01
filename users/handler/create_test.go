package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/common/validations"
	"github.com/ben-toogood/kite/users"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	h := testHandler(t)

	t.Run("MissingFirstName", func(t *testing.T) {
		rsp, err := h.Create(context.TODO(), &users.CreateRequest{
			LastName: "Doe",
			Email:    "johndoe@gmail.com",
		})
		assert.Nil(t, rsp)

		var found bool
		for _, e := range validations.ExtractErrors(err) {
			if e.Field == "first_name" {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("MissingLastName", func(t *testing.T) {
		rsp, err := h.Create(context.TODO(), &users.CreateRequest{
			FirstName: "John",
			Email:     "johndoe@gmail.com",
		})
		assert.Nil(t, rsp)

		var found bool
		for _, e := range validations.ExtractErrors(err) {
			if e.Field == "last_name" {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("MissingEmail", func(t *testing.T) {
		rsp, err := h.Create(context.TODO(), &users.CreateRequest{
			FirstName: "John",
			LastName:  "Doe",
		})
		assert.Nil(t, rsp)

		var found bool
		for _, e := range validations.ExtractErrors(err) {
			if e.Field == "email" {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("DuplicateEmail", func(t *testing.T) {
		req := &users.CreateRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     uuid.New().String() + "@gmail.com",
		}
		_, err := h.Create(context.TODO(), req)
		assert.NoError(t, err)

		_, err = h.Create(context.TODO(), req)
		assert.Equal(t, database.ErrDuplicate, err)
	})

	t.Run("Valid", func(t *testing.T) {
		req := &users.CreateRequest{
			FirstName: "John",
			LastName:  "Doe",
			Email:     uuid.New().String() + "@gmail.com",
		}
		rsp, err := h.Create(context.TODO(), req)
		assert.NoError(t, err)
		assert.NotNil(t, rsp)
		if rsp == nil {
			return
		}
		assert.NotNil(t, rsp.User)
		if rsp.User == nil {
			return
		}
		assert.Contains(t, rsp.User.Id, "usr_")
		assert.NotNil(t, rsp.User.CreatedAt)
		assert.NotNil(t, rsp.User.UpdatedAt)
		assert.Equal(t, req.FirstName, rsp.User.FirstName)
		assert.Equal(t, req.LastName, rsp.User.LastName)
		assert.Equal(t, req.Email, rsp.User.Email)

	})
}
