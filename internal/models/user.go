package models

import (
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(20);not null" form:"name" binding:"required"`
	Email   string `gorm:"type:varchar(50);unique" form:"email" binding:"required,email"`
	Phone   string `gorm:"type:varchar(13);unique;not null" form:"phone" binding:"required"`
	Init    bool
	Version optimisticlock.Version
}
