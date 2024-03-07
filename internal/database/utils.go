package database

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func buildGormQuery(db *gorm.DB, queries map[string]string) *gorm.DB {
	query := db
	for field, value := range queries {
		if value == "" {
			continue
		}
		query = query.Where(fmt.Sprintf("%s=?", field), value)
	}
	return query
}

func generateUniqueCode() string {
	return uuid.New().String()
}
