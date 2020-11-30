package model

import (
	"time"

	"github.com/ben-toogood/kite/comments"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID           string
	ResourceType comments.ResourceType
	ResourceID   string
	AuthorID     string
	Message      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (c *Comment) BeforeCreate(scope *gorm.DB) error {
	c.ID = "cmt_" + ksuid.New().String()
	return nil
}
