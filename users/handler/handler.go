package handler

import (
	"github.com/lileio/pubsub/v2"
	"gorm.io/gorm"
)

// Users implements the users handler interface
type Users struct {
	DB     *gorm.DB
	PubSub *pubsub.Client
}
