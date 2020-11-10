package csv

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type FileConfiguration struct {
	gorm.Model
	ID             int
	Name           string
	FilePattern    string
	ColumnMappings []*ColumnMapping //`gorm:"foreignKey:file_configuration_id;references:id"`
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
}

func (FileConfiguration) TableName() string {
	return "csv_file_configurations"
}

func (fp *FileConfiguration) ParseRow(row RowData) (Record, error) {
	record := make(map[string]interface{})
	fmt.Println("PARSING ROW")
	fmt.Println(len(fp.ColumnMappings))

	for _, col := range fp.ColumnMappings {
		if col.Ordinal-1 > len(record) {
			return nil, fmt.Errorf("Column mapping ordinal outside record index")
		}
		key := col.Name
		val, err := col.Get(row[col.Ordinal-1])
		fmt.Println(col.Name)
		fmt.Println(val.(string))

		if err != nil {
			return nil, err
		}

		record[key] = val
	}

	return record, nil
}
