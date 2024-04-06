package models

import "gorm.io/gorm/clause"

func ImportData(data any) (err error) {
	tx := db
	if err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(data).Error; err != nil {
		return
	}
	return
}
