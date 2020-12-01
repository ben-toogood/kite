package model

import (
	"time"

	"github.com/ben-toogood/kite/followers"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID            string
	FollowerType  followers.ResourceType
	FollowerID    string
	FollowingType followers.ResourceType
	FollowingID   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (f *Follow) BeforeCreate(scope *gorm.DB) error {
	f.ID = "fol_" + ksuid.New().String()
	return nil
}
