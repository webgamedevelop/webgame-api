package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	if err = createDatabase(); err != nil {
		return err
	}
	if db, err = gorm.Open(mysql.Open(dsn())); err != nil {
		return err
	}
	return nil
}
