package csv

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	db.Migrator().DropTable(
		&FileConfiguration{},
		&ColumnMapping{},
	)

	err := db.AutoMigrate(
		&FileConfiguration{},
		&ColumnMapping{},
	)

	if err != nil {
		return err
	}

	return seedData(db)
}

func seedData(db *gorm.DB) error {
	config := &FileConfiguration{
		Name:        "Discover All Available",
		FilePattern: "Discover-AllAvailable-[0-9]+.csv",
		ColumnMappings: []ColumnMapping{
			ColumnMapping{
				Name:    "Transaction Date",
				Ordinal: 1,
				Type:    "timestamp",
				Args:    map[string]interface{}{"timestamp_format": "01/02/2006"},
			},
			ColumnMapping{
				Name:    "Posted Date",
				Ordinal: 2,
				Type:    "timestamp",
				Args:    map[string]interface{}{"timestamp_format": "01/02/2006"},
			},
			ColumnMapping{
				Name:    "Description",
				Ordinal: 3,
				Type:    "string",
			},
			ColumnMapping{
				Name:    "Amount",
				Ordinal: 4,
				Type:    "float64",
			},
			ColumnMapping{
				Name:    "Category",
				Ordinal: 5,
				Type:    "string",
			},
		},
	}

	result := db.Create(&config)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
