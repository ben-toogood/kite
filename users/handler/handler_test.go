package handler

import (
	"context"
	"testing"

	"github.com/ben-toogood/kite/common/database"
	"github.com/ben-toogood/kite/users/model"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Users {
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.User{})
	assert.NoErrorf(t, err, "Error migrating database")

	psc := &pubsub.Client{
		ServiceName: "users",
		// Provider:    //counterfitter provider,
		Middleware: defaults.Middleware,
	}

	return &Users{DB: db, PubSub: psc}
}
