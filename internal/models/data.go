package models

import (
	"gorm.io/gorm/clause"
)

// initialization data
var (
	IngressClasses = []*IngressClass{
		{
			Name:      "nginx-ingress",
			ClassName: "nginx",
		},
	}
	ResourceSpecs = []*ResourceSpec{
		{
			Name:   "small",
			Cpu:    "50m",
			Memory: "128Mi",
		}, {
			Name:   "medium",
			Cpu:    "100m",
			Memory: "256Mi",
		}, {
			Name:   "large",
			Cpu:    "200m",
			Memory: "512Mi",
		},
	}
)

// Initialize import initialization data
func Initialize() error {
	var err error
	tx := db.Begin()

	// rollback when panic or err
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return err
	}

	if err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(IngressClasses).Error; err != nil {
		return err
	}

	if err = tx.Clauses(clause.OnConflict{DoNothing: true}).Create(ResourceSpecs).Error; err != nil {
		return err
	}

	if err = tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
