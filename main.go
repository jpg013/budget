package budget

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jpg013/budget/csv"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Activity struct {
// 	gorm.Model
// 	ActivityID      int
// 	TransactionDate time.Time
// 	PostedDate      time.Time
// 	Description     string
// 	Amount          float32
// 	Category        string
// }

// func main() {
// 	dsn := "user=dev password=dev dbname=dev port=5432 sslmode=disable host=localhost"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err)
// 	}

// 	csvFileParser, err := csv.MakeFileParser(db)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(csvFileParser)

// 	// Load the file configuration

// 	// fmt.Println(db)

// 	// f, err := os.Open("./Discover-2020-YearToDateSummary.csv")
// 	// if err != nil {
// 	// 	log.Fatal("Unable to read input file ", err)
// 	// }
// 	// defer f.Close()

// 	// csvReader := csv.NewReader(f)
// 	// records, err := csvReader.ReadAll()
// 	// if err != nil {
// 	// 	log.Fatal("Unable to parse file as CSV", err)
// 	// }

// 	// // Remove the first record (these are the column names)
// 	// records = records[1:]
// 	// // rowNum := 0

// 	// for _, r := range records {
// 	// 	s := strings.Split(r[0], "/")
// 	// 	s1 := fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
// 	// 	td, err := time.Parse("2006-01-02", s1)

// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	s = strings.Split(r[1], "/")
// 	// 	s1 = fmt.Sprintf("%s-%s-%s", s[2], s[0], s[1])
// 	// 	pd, err := time.Parse("2006-01-02", s1)

// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}

// 	// 	fmt.Println("================")
// 	// 	fmt.Println(r[0])
// 	// 	fmt.Println(r[1])
// 	// 	fmt.Println(td)
// 	// 	fmt.Println(pd)
// 	// 	// pd, _ := time.Parse(record[1], "11-01-2012")
// 	// 	// // activity = &Activity{
// 	// 	// // 	TransactionDate: time.Parse(record[0], '10/19/2020')
// 	// 	// // }

// 	// 	// // date := "1999-12-31"
// 	// 	// // t, _ := time.Parse(layoutISO, date)
// 	// 	// fmt.Println(err)
// 	// 	// fmt.Println(td)
// 	// 	// fmt.Println(pd)
// 	// }
// }
