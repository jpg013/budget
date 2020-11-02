package models

import (
	"time"

	"gorm.io/gorm"
)

type ActivitySource struct {
	gorm.Model
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
