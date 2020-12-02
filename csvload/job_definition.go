package csvload

import (
	"fmt"
)

type JSONB map[string]interface{}

type ColumnType string

const (
	IntColumn       ColumnType = "int"
	Float64Column              = "float64"
	StrColumn                  = "string"
	TimestampColumn            = "timestamp"
)

type JobType string

const (
	CSVFileLoad JobType = "csv_file_load"
)

// JobDefinition represents a file load job definition
type JobDefinition struct {
	Name        string
	Type        JobType
	FilePattern string
	Columns     []CSVColumn
}

type CSVColumn struct {
	Name    string     `json:"name"`
	Key     string     `json:"key"`
	Ordinal int        `json:"ordinal"`
	Type    ColumnType `json:"type"`
	Args    JSONB      `json:"args"`
}

type parseableJobDefinition struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	FilePattern string `json:"file_pattern"`
	Columns     []parseableCSVColumn
}

type parseableCSVColumn struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	Ordinal int    `json:"ordinal"`
	Type    string `json:"type"`
	Args    JSONB  `json:"args"`
}

func (p *parseableCSVColumn) normalize() (c CSVColumn, err error) {
	ct, err := parseColumnType(p.Type)

	if err != nil {
		return
	}

	return CSVColumn{
		Name:    p.Name,
		Key:     p.Key,
		Ordinal: p.Ordinal,
		Type:    ct,
		Args:    p.Args,
	}, err
}

func (p *parseableJobDefinition) normalize() (def JobDefinition, err error) {
	jt, err := parseJobtype(p.Type)

	if err != nil {
		return
	}

	cols := make([]CSVColumn, len(p.Columns))

	for idx, c := range p.Columns {
		col, err := c.normalize()
		if err != nil {
			return def, err
		}
		cols[idx] = col
	}

	return JobDefinition{
		Name:        p.Name,
		Type:        jt,
		FilePattern: p.FilePattern,
		Columns:     cols,
	}, err
}

func parseColumnType(s string) (ColumnType, error) {
	switch s {
	case "int":
		return IntColumn, nil
	case "float64":
		return Float64Column, nil
	case "string":
		return StrColumn, nil
	case "timestamp":
		return TimestampColumn, nil
	default:
		return "", fmt.Errorf("invalid column type \"%s\"", s)
	}
}

func parseJobtype(s string) (JobType, error) {
	switch s {
	case "csv_file_load":
		return CSVFileLoad, nil
	default:
		return "", fmt.Errorf("invalid job type \"%s\"", s)
	}
}
