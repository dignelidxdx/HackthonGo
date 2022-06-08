package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm(dataSource string) (*gorm.DB, error) {

	var err error

	gorm, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		log.Println("database Configured")
		return nil, err
	}

	return gorm, nil
}
