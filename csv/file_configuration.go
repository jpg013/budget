package csv

import (
	"gorm.io/gorm"
)

type FileConfiguration struct {
	gorm.Model
	ID               int
	Name             string
	FilePattern      string
	ColumnMappings   []ColumnMapping `gorm:"foreignKey:file_configuration_id;references:id"`
	ActivitySourceID int             `gorm:"not null"`
}

// TableName overrides the table name used by User to `profiles`
func (FileConfiguration) TableName() string {
	return "csv_file_configurations"
}
