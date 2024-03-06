package database

import (
	"fmt"
	"gorm.io/gorm"
)

func buildGormQuery(db *gorm.DB, queries map[string]string) *gorm.DB {
	for field, value := range queries {
		if value == "" {
			continue
		}
		db.Where(fmt.Sprintf("%s = ?", field), value)
	}
	return db
}
