package handler

import (
	"crypto/rsa"

	"github.com/lileio/pubsub/v2"
	"gorm.io/gorm"
)

// Auth implements the auth handler interface
type Auth struct {
	DB            *gorm.DB
	PubSub        *pubsub.Client
	JWTPrivateKey *rsa.PrivateKey
	JWTPublicKey  *rsa.PublicKey
}
