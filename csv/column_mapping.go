package csv

import (
	"gorm.io/gorm"
)

const (
	IntColumn       ColumnType = "int"
	Float64Column              = "float64"
	StrColumn                  = "string"
	TimestampColumn            = "timestamp"
)

type ColumnMapping struct {
	gorm.Model
	Name                string
	Ordinal             int `gorm:"UNIQUE_INDEX:uniqueindex;type:int;not null"`
	Type                ColumnType
	Args                JSONB `sql:"type:jsonb"`
	FileConfigurationID int   `gorm:"UNIQUE_INDEX:uniqueindex;type:int;not null"`
}

// TableName overrides the table name used by User to `profiles`
func (ColumnMapping) TableName() string {
	return "csv_column_mappings"
}
