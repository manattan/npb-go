package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	// Ref. https://gorm.io/ja_JP/docs/connecting_to_the_database.html
	dsn := "root@tcp(127.0.0.1:3306)/npb_v1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	return db, nil
}
