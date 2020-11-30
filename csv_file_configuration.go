package budget

import (
	"gorm.io/gorm"
)

type CSVFileConfiguration struct {
	gorm.Model
	Name       string      `gorm:"unique"`
	CSVColumns []CSVColumn `gorm:"foreignKey:file_configuration_id;references:id"`
}
