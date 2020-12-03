package handler

import (
	"crypto/rsa"
	"time"

	"github.com/ben-toogood/kite/users"
	"github.com/lileio/pubsub/v2"
	sg "github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gorm.io/gorm"
)

const issuer = "kite"

var (
	accessTokenTTL  = time.Minute * 10
	refreshTokenTTL = time.Hour * 24 * 7
)

type sendgridClient interface {
	Send(*mail.SGMailV3) (*sg.Response, error)
}

// Auth implements the auth handler interface
type Auth struct {
	Users      users.UsersServiceClient
	DB         *gorm.DB
	PubSub     *pubsub.Client
	PrivateKey *rsa.PrivateKey
	Sendgrid   sendgridClient
}
