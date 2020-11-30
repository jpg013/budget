package budget

import "gorm.io/gorm"

type CSVFileMapping struct {
	gorm.Model
	Name        string
	Pattern     string `gorm:"unique"`
	CSVConfigID uint
}
