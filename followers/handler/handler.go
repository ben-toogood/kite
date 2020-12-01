package handler

import (
	"gorm.io/gorm"
)

// Followers implements the followers handler interface
type Followers struct {
	DB *gorm.DB
}
