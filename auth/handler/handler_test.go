package handler

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"strings"
	"testing"

	"github.com/ben-toogood/kite/auth/model"
	"github.com/ben-toogood/kite/common/database"
	"github.com/lileio/pubsub/v2"
	"github.com/lileio/pubsub/v2/middleware/defaults"
	sg "github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/assert"
)

type sendgridMock struct {
	Messages []string
}

func (s *sendgridMock) Send(m *mail.SGMailV3) (*sg.Response, error) {
	s.Messages = append(s.Messages, m.Content[0].Value)
	return nil, nil
}

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
		Sendgrid:      &sendgridMock{Messages: []string{}},
	}
}

// get the jwt from the message
func getJWTFromTestHandlerEmail(t *testing.T, h *Auth) (string, error) {
	msg := h.Sendgrid.(*sendgridMock).Messages[0]
	if msg == "" {
		return "", errors.New("Message not found")
	}
	cp1 := strings.Split(msg, "code=")
	if len(cp1) != 2 {
		return "", errors.New("Message not in correct format")
	}
	cp2 := strings.Split(cp1[1], " ")
	if len(cp2) < 1 {
		return "", errors.New("Message not in correct format")
	}

	return cp2[0], nil
}
