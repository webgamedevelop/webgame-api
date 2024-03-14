package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() error {
	var err error
	if err = createDatabase(); err != nil {
		return err
	}

	if debugLogLevel {
		if db, err = gorm.Open(mysql.Open(dsn()), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}); err != nil {
			return err
		}
		return nil
	}

	if db, err = gorm.Open(mysql.Open(dsn())); err != nil {
		return err
	}
	return nil
}
