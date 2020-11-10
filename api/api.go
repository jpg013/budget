package api

import (

	// Importing go sql driver

	"fmt"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jpg013/budget"
	"github.com/jpg013/budget/config"
	"github.com/jpg013/budget/csv"
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

func DefineDiscoverAllAvailableCSV(db *gorm.DB) {
	// config := &csv.FileConfiguration{
	// 	Name:        "Discover All Available",
	// 	FilePattern: "Discover-AllAvailable-[0-9]+.csv",
	// 	ColumnMappings: []*csv.ColumnMapping{
	// 		&csv.ColumnMapping{
	// 			Name:    "Transaction Date",
	// 			Ordinal: 1,
	// 			Type:    "timestamp",
	// 			Args:    map[string]interface{}{"timestamp_format": "01-02-2006"},
	// 		},
	// 		&csv.ColumnMapping{
	// 			Name:    "Posted Date",
	// 			Ordinal: 2,
	// 			Type:    "timestamp",
	// 			Args:    map[string]interface{}{"timestamp_format": "01-02-2006"},
	// 		},
	// 		&csv.ColumnMapping{
	// 			Name:    "Description",
	// 			Ordinal: 3,
	// 			Type:    "string",
	// 		},
	// 		&csv.ColumnMapping{
	// 			Name:    "Amount",
	// 			Ordinal: 4,
	// 			Type:    "float64",
	// 		},
	// 		&csv.ColumnMapping{
	// 			Name:    "Category",
	// 			Ordinal: 5,
	// 			Type:    "string",
	// 		},
	// 	},
	// }

	// result := db.Create(&config)
	// fmt.Println(result)
}

func Start(cfg config.Configuration) error {
	dsn := makePostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	fileWatcher := budget.NewFileWatcher()
	csvManager, err := csv.NewManager(db)

	if err != nil {
		return err
	}

	go func() {
		err := fileWatcher.WatchDirs("./tmp")

		if err != nil {
			panic(err)
		}

		for {
			evt := <-fileWatcher.Data

			if evt.Op == fsnotify.Create {
				records, err := csvManager.ParseFile(evt.Name)

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(records)
				}
			}
		}
	}()

	fmt.Println(csvManager)

	time.Sleep(1000 * time.Second)
	return nil
}
