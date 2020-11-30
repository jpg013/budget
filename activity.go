package budget

import "time"

type Activity struct {
	TransactionDate  *time.Time
	PostedDate       *time.Time
	Description      string
	Amount           float64
	Category         string
	Code             string `gorm:"index:activity_code,unique"`
	ActivitySourceID int    `gorm:"UNIQUE_INDEX:uniqueindex;type:int;not null"`
}
