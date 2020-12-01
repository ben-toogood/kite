package handler

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	"github.com/stretchr/testify/assert"
)

func testHandler(t *testing.T) *Auth {
	// connect to database
	db, err := database.GetDB(context.TODO())
	assert.NoErrorf(t, err, "Error connecting to database")
	err = db.AutoMigrate(&model.Token{})
	assert.NoErrorf(t, err, "Error migrating database")

	// connect to pubsub
	psc := &pubsub.Client{
		ServiceName: "Auth",
		Middleware:  defaults.Middleware,
	}

	// generate JWT private / public keys
	reader := rand.Reader
	bitSize := 2048
	key, err := rsa.GenerateKey(reader, bitSize)
	assert.NoErrorf(t, err, "Error generating public key")

	return &Auth{
		DB:            db,
		PubSub:        psc,
		JWTPrivateKey: key,
		JWTPublicKey:  &key.PublicKey,
	}
}
