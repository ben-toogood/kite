package model

import (
	"time"

	"github.com/ben-toogood/kite/likes"
)

type Like struct {
	ResourceID   string             `gorm:"index:usr_like,unique"`
	ResourceType likes.ResourceType `gorm:"index:usr_like,unique"`
	UserID       string             `gorm:"index:usr_like,unique"`
	CreatedAt    time.Time
}
