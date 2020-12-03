package model

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Post struct {
	ID          string
	AuthorID    string
	ImageID     string `pb:"ignore=true"`
	Description string
	CreatedAt   time.Time
}

func (p *Post) BeforeCreate(scope *gorm.DB) error {
	p.ID = "pos_" + ksuid.New().String()
	return nil
}
