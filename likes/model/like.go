package model

import (
	"time"

	"github.com/ben-toogood/kite/likes"
)

type Like struct {
	ResourceID   string
	ResourceType likes.ResourceType
	UserID       string
	CreatedAt    time.Time
}
