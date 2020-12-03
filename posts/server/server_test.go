package server

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/posts/model"
	"github.com/stretchr/testify/assert"
)

func testServer(t *testing.T) *Posts {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.Post{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Posts{DB: db}
}
