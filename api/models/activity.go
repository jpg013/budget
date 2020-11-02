package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	ID              int
	TransactionDate time.Time
	PostedDate      time.Time
	Description     string
	Amount          float32
	Category        string
	Code            string
	SourceID        int
}
