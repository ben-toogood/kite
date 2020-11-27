package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users/model"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Users {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.User{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Users{DB: db}
}
