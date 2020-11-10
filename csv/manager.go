package csv

import (
	encodingcsv "encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"gorm.io/gorm"
)

type Manager struct {
	fileConfigs []*FileConfiguration
}

func (m *Manager) getConfigForFile(filename string) (*FileConfiguration, error) {
	for _, conf := range m.fileConfigs {
		if conf.FilePattern == "" {
			continue
		}

		match, err := regexp.MatchString(conf.FilePattern, filename)
		if err != nil {
			return nil, err
		}

		if match {
			return conf, nil
		}
	}

	return nil, nil
}

func (m *Manager) ParseFile(file string) ([]Record, error) {
	ext := filepath.Ext(file)

	if ext != ".csv" {
		return nil, fmt.Errorf("file of type \"%s\" is not a valid CSV", ext)
	}

	fileConfig, err := m.getConfigForFile(file)

	if err != nil {
		return nil, fmt.Errorf("Could not get file configration for file %s: %v", file, err)
	}

	// Read file
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("Unable to read input file %v", err)
	}

	defer f.Close()

	csvReader := encodingcsv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("Unable to parse file as CSV input file %v", err)
	}

	data := make([]Record, len(records)-1)

	for idx, r := range records[1:] {
		datum, err := fileConfig.ParseRow(r)

		if err != nil {
			return nil, fmt.Errorf("Unable to parse file as CSV input file %v", err)
		}

		data[idx] = datum
	}

	return data, nil
}

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

func NewManager(db *gorm.DB) (m *Manager, err error) {
	// result := db.AutoMigrate(&User{}, &Order{})

	// fmt.Println(result)

	// user := User{
	// 	Username: "Justin",
	// 	Orders: []Order{
	// 		Order{
	// 			Price: 12.42,
	// 		},
	// 	},
	// }

	// result := db.Create(&user)
	// fmt.Println(result)
	var user []User
	result := db.Table("users").Joins("orders").Find(&user, "username='Justin'")
	fmt.Println(result)
	fmt.Println(user)
	os.Exit(0)

	// m = &Manager{}
	// // Load all csv file configurations
	// result := db.Joins("ColumnMappings").Find(&m.fileConfigs)
	// fmt.Println(result)
	// fmt.Println(m.fileConfigs[0].Name)
	// panic("no")
	// if result.Error() != "" {
	// 	fmt.Println(result.Error())
	// 	return m, err
	// }
	// fmt.Println("Bitch please")
	// fmt.Println(result)

	return m, nil
}
