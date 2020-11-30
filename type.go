package budget

import (
	"database/sql/driver"
	"encoding/json"
)

type CSVRowData []string

type CSVRecord map[string]interface{}

type CSVColumnType string

type JSONB map[string]interface{}

type ColumnType string

type RowData []string

type Record map[string]interface{}

// type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	s, ok := value.(string)

	if !ok {
		return nil
	}

	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return err
	}

	return nil
}
