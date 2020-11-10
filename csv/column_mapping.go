package csv

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type ColumnType string

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}

const (
	IntColumn       ColumnType = "int"
	Float64Column              = "float64"
	StrColumn                  = "string"
	TimestampColumn            = "timestamp"
)

type ColumnMapping struct {
	gorm.Model
	Name    string
	Ordinal int
	Type    ColumnType
	// Args                JSONB `sql:"type:jsonb"`
	FileConfigurationID int
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	DeletedAt           *time.Time
}

func (cm *ColumnMapping) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AFTER FIND??")
	return err
}

func (ColumnMapping) TableName() string {
	return "csv_column_mappings"
}

func (cm *ColumnMapping) Get(v interface{}) (interface{}, error) {
	switch cm.Type {
	case Float64Column:
		return cm.parseString(v)
	case StrColumn:
		return cm.parseString(v)
	case TimestampColumn:
		return cm.parseTimestamp(v)
	default:
		return nil, errors.New("invalid column mapping type")
	}
}

func (cm *ColumnMapping) parseString(v interface{}) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

func (cm *ColumnMapping) parseFloat(v interface{}) (float64, error) {
	s, err := cm.parseString(v)

	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(s, 64)
}

func (cm *ColumnMapping) parseTimestamp(v interface{}) (t time.Time, err error) {
	s, err := cm.parseString(v)

	if err != nil {
		return t, err
	}

	// Parse timestamp format from args
	// format, ok := cm.Args["format_timestamp"].(string)

	// if !ok {
	format := "2006-01-02"
	// }

	t, err = time.Parse(format, s)

	return t, err
}

// func MakeColumnMapping(s string, n string, o int, v interface{}) (*ColumnMapping, error) {
// 	columnType, err := ParseColumnType(s)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ColumnMapping{
// 		Name:    n,
// 		Ordinal: 0,
// 		Type:    columnType,
// 	}, nil
// }

// func ParseColumnType(s string) (ColumnType, error) {
// 	switch s {
// 	case "int":
// 		return IntColumn, nil
// 	case "float64":
// 		return Float64Column, nil
// 	case "string":
// 		return StrColumn, nil
// 	case "timestamp":
// 		return TimestampColumn, nil
// 	default:
// 		return "", errors.New("Invalid column type")
// 	}
// }
