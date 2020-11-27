package model

import (
	"time"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(scope *gorm.DB) error {
	u.ID = "usr_" + ksuid.New().String()
	return nil
}
