package handler

import (
	"gorm.io/gorm"
)

// Users implements the users handler interface
type Users struct {
	DB *gorm.DB
}
