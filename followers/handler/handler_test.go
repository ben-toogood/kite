package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/followers/model"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Followers {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.Follow{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Followers{DB: db}
}
