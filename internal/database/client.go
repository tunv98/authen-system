package database

import (
	"authen-system/internal/config"
	"authen-system/internal/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func ProvideSQL(config config.MySQL) (*gorm.DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.PassWord,
		config.HostPort,
		config.DatabaseName,
	)
	db, err := gorm.Open(mysql.Open(connectString))
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&models.User{}, &models.Voucher{}, &models.Campaign{}); err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.ConnectionMaxLifetime))
	return db, nil
}
