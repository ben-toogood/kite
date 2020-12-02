package handler

import (
	"gorm.io/gorm"
)

// Likes implements the likes handler interface
type Likes struct {
	DB *gorm.DB
}
