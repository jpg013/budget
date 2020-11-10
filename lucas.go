package budget

// import (
// 	"fmt"

// 	"github.com/jpg013/budget/csv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func main() {
// 	config := &csv.FileConfiguration{
// 		Name:        "Discover All Available",
// 		FilePattern: "Discover-AllAvailable-.csv",
// 		ColumnMappings: []*csv.ColumnMapping{
// 			&csv.ColumnMapping{
// 				Name:    "Transaction Date",
// 				Ordinal: 1,
// 				Type:    "timestamp",
// 				Args:    map[string]interface{}{"timestamp_format": "01-02-2006"},
// 			},
// 			&csv.ColumnMapping{
// 				Name:    "Posted Date",
// 				Ordinal: 2,
// 				Type:    "timestamp",
// 				Args:    map[string]interface{}{"timestamp_format": "01-02-2006"},
// 			},
// 			&csv.ColumnMapping{
// 				Name:    "Description",
// 				Ordinal: 3,
// 				Type:    "string",
// 			},
// 			&csv.ColumnMapping{
// 				Name:    "Amount",
// 				Ordinal: 4,
// 				Type:    "float64",
// 			},
// 			&csv.ColumnMapping{
// 				Name:    "Category",
// 				Ordinal: 5,
// 				Type:    "string",
// 			},
// 		},
// 	}

// 	dsn := "user=dev password=dev dbname=dev port=5432 sslmode=disable host=localhost"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	result := db.Create(&config)
// 	fmt.Println(result)
// }
