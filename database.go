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

	if err != nil {
		return db, err
	}

	err = migrate(db)

	if err != nil {
		return db, err
	}

	err = seedData(db)

	if err != nil {
		return db, err
	}

	return db, err
}

func migrate(db *gorm.DB) error {
	// db.Migrator().DropTable(
	// 	&csv.Configuration{},
	// 	&csv.Column{},
	// 	&ActivitySource{},
	// 	&CSVFileMapping{},
	// 	&Activity{},
	// )
	return nil
	// return db.AutoMigrate(
	// 	&csv.Configuration{},
	// 	&csv.Column{},
	// 	&ActivitySource{},
	// 	&CSVFileMapping{},
	// 	&Activity{},
	// )
}

func seedData(db *gorm.DB) error {
	// configuration := csv.Configuration{
	// 	Name: "Discover All Available",
	// 	Columns: []csv.Column{
	// 		csv.Column{
	// 			Name:        "Transaction Date",
	// 			Key:         "transaction_date",
	// 			Ordinal:     1,
	// 			Type:        "timestamp",
	// 			Args:        map[string]interface{}{"timestamp_format": "01/02/2006"},
	// 			IsKeyColumn: true,
	// 		},
	// 		csv.Column{
	// 			Name:        "Posted Date",
	// 			Key:         "posted_date",
	// 			Ordinal:     2,
	// 			Type:        "timestamp",
	// 			Args:        map[string]interface{}{"timestamp_format": "01/02/2006"},
	// 			IsKeyColumn: true,
	// 		},
	// 		csv.Column{
	// 			Name:        "Description",
	// 			Key:         "descriptions",
	// 			Ordinal:     3,
	// 			Type:        "string",
	// 			IsKeyColumn: true,
	// 		},
	// 		csv.Column{
	// 			Name:        "Amount",
	// 			Key:         "amount",
	// 			Ordinal:     4,
	// 			Type:        "float64",
	// 			IsKeyColumn: true,
	// 		},
	// 		csv.Column{
	// 			Name:        "Category",
	// 			Key:         "category",
	// 			Ordinal:     5,
	// 			Type:        "string",
	// 			IsKeyColumn: true,
	// 		},
	// 	},
	// }

	// result := db.Create(&configuration)

	// if result.Error != nil {
	// 	return result.Error
	// }

	// activitySource := ActivitySource{
	// 	Name: "discover",
	// }

	// result = db.Create(&activitySource)

	// if result.Error != nil {
	// 	return result.Error
	// }

	// mapping := CSVFileMapping{
	// 	Name:        "Discover All Available",
	// 	Pattern:     "Discover-AllAvailable-[0-9]+.csv",
	// 	CSVConfigID: configuration.ID,
	// }

	// if result.Error != nil {
	// 	return result.Error
	// }

	return nil
}
