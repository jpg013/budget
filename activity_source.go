package budget

import "gorm.io/gorm"

type ActivitySource struct {
	gorm.Model
	Name string `gorm:"index:activity_source_name,unique"`
	// Activities []Activity `gorm:"foreignKey:activity_source_id;references:id"`
}
