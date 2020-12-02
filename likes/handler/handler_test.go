package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/likes/model"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Likes {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.Like{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Likes{DB: db}
}
