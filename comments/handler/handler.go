package handler

import (
	"gorm.io/gorm"
)

// Comments implements the comments handler interface
type Comments struct {
	DB *gorm.DB
}
