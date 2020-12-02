package server

import (
	"context"
	"log"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users/model"
	"github.com/stretchr/testify/assert"
)

func testServer(t *testing.T) *Users {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.User{})
	assert.NoErrorf(t, err, "Error migrating database")
	return &Users{}
}
