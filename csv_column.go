package budget

import (
	"gorm.io/gorm"
)

const (
	IntColumn       ColumnType = "int"
	Float64Column              = "float64"
	StrColumn                  = "string"
	TimestampColumn            = "timestamp"
)

type CSVColumn struct {
	gorm.Model
	Name            string `gorm:"type:varchar;not null"`
	Key             string `gorm:"type:varchar;not null"`
	Ordinal         int    `gorm:"UNIQUE_INDEX:uniqueindex;type:int;not null"`
	Type            ColumnType
	Args            JSONB `sql:"type:jsonb"`
	ConfigurationID int   `gorm:"UNIQUE_INDEX:uniqueindex;type:int;not null"`
	IsKeyColumn     bool
}
