package budget

import (
	"fmt"
	"os"

	"github.com/jpg013/budget/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func makePostgresDSN(cfg config.Configuration) string {
	pass := os.Getenv("POSTGRES_PASSWORD")

	// Yea this is for dev only
	if pass == "" {
		pass = "dev"
	}

	return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable host=%s",
		cfg.Postgres.Username,
		pass,
		cfg.Postgres.Database,
		cfg.Postgres.Port,
		cfg.Postgres.Host,
	)
}

func CreateDB(cfg config.Configuration) (*gorm.DB, error) {
	dsn := makePostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

	// if err != nil {
	// 	return nil, err
	// }

	// db.Migrator().DropTable(
	// 	&models.CSVFileConfiguration{},
	// 	&models.CSVColumnMapping{},
	// 	&models.ActivitySource{},
	// 	&models.Activity{},
	// )

	// err = autoMigration(db)

	// if err != nil {
	// 	return nil, err
	// }

	// err = defineDiscoverAllAvailableCSV(db)

	// if err != nil {
	// 	return nil, err
	// }

	// return db, err
}

// func autoMigration(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		&models.CSVFileConfiguration{},
// 		&models.CSVColumnMapping{},
// 		&models.ActivitySource{},
// 		&models.Activity{},
// 	)
// }

// func defineDiscoverAllAvailableCSV(db *gorm.DB) error {
// 	config := &CSVFileConfiguration{
// 		Name:        "Discover All Available",
// 		FilePattern: "Discover-AllAvailable-[0-9]+.csv",
// 		ColumnMappings: []CSVColumnMapping{
// 			CSVColumnMapping{
// 				Name:    "Transaction Date",
// 				Ordinal: 1,
// 				Type:    "timestamp",
// 				Args:    map[string]interface{}{"timestamp_format": "01/02/2006"},
// 			},
// 			CSVColumnMapping{
// 				Name:    "Posted Date",
// 				Ordinal: 2,
// 				Type:    "timestamp",
// 				Args:    map[string]interface{}{"timestamp_format": "01/02/2006"},
// 			},
// 			CSVColumnMapping{
// 				Name:    "Description",
// 				Ordinal: 3,
// 				Type:    "string",
// 			},
// 			CSVColumnMapping{
// 				Name:    "Amount",
// 				Ordinal: 4,
// 				Type:    "float64",
// 			},
// 			CSVColumnMapping{
// 				Name:    "Category",
// 				Ordinal: 5,
// 				Type:    "string",
// 			},
// 		},
// 	}

// 	result := db.Create(&config)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return nil
// }
