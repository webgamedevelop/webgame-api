package models

import (
	"errors"
	"fmt"
	"runtime/debug"

	"gorm.io/gorm/clause"
)

func Delete(id uint, dest any, fn func() error) (err error) {
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

	if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(dest, id).Error; err != nil {
		return
	}

	if err = tx.Delete(dest).Error; err != nil {
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
