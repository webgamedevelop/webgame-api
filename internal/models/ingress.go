package models

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// IngressClass k8s ingress class
type IngressClass struct {
	gorm.Model
	// Display name
	Name string `gorm:"type:varchar(50);uniqueIndex:idx_name,priority:1;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	// K8S ingress class resource name
	// Cannot be updated
	ClassName string `gorm:"type:varchar(20);uniqueIndex:idx_class_name,priority:1;not null" binding:"required,max=20" form:"className" json:"className,omitempty"`
	// Imported flag
	Imported bool `form:"imported" json:"imported,omitempty"`
	// Synced flag
	Synced bool  `form:"synced" json:"synced,omitempty"`
	DelAt  int64 `gorm:"uniqueIndex:idx_name,priority:2;uniqueIndex:idx_class_name,priority:2;not null" json:"delAt,omitempty"`
}

func (i *IngressClass) BeforeDelete(tx *gorm.DB) (err error) {
	if i.Imported {
		return fmt.Errorf("cat not delete imported ingress class, name: %s, className: %s", i.Name, i.ClassName)
	}
	i.DelAt = time.Now().UnixMicro()
	if err = tx.Updates(i).Error; err != nil {
		return
	}
	return
}

func (i *IngressClass) Update(fn func() error) (err error) {
	tx := db.Begin()
	// rollback when panic or err
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("panic in transaction: %s", r)
			}
			err = errors.Join(err, fmt.Errorf("stack: \n%s\n", stack))
			tx.Rollback()
			return
		}
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if err = tx.Error; err != nil {
		return
	}

	name := i.Name
	if err = tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).First(i).Error; err != nil {
		return
	}

	i.Name = name
	if err = tx.Updates(i).Error; err != nil {
		return
	}

	if err = fn(); err != nil {
		return
	}

	if err = tx.Commit().Error; err != nil {
		return
	}

	return
}

func (i *IngressClass) Detail() (err error) {
	if err = db.First(i).Error; err != nil {
		return
	}
	return
}
