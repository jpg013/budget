package models

import (
	"time"

	"gorm.io/gorm"
)

type ActivityFile struct {
	gorm.Model
	ID          int
	Name        string
	Type        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	ProcessedAt time.Time
	Extension   string
}
