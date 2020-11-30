package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/comments/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Comments {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.Comment{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Comments{DB: db}
}
