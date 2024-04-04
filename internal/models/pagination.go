package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Paginator struct {
	Page     int    `binding:"required,min=1" form:"page" json:"page"`         // page
	PageSize int    `binding:"required,min=1" form:"pageSize" json:"pageSize"` // page size
	Column   string `form:"column" json:"column"`                              // column name to order by
	Desc     bool   `form:"desc" json:"desc"`                                  // desc
}

func (p *Paginator) Offset() int {
	return p.PageSize * (p.Page - 1)
}

func (p *Paginator) SetDefault() {
	if p.Column == "" {
		p.Column = "id"
	}
}

func List(dest interface{}, paginator *Paginator, conditions func(db *gorm.DB) (*gorm.DB, error)) error {
	var (
		tx  *gorm.DB
		err error
	)

	tx = db
	if conditions != nil {
		if tx, err = conditions(tx); err != nil {
			return err
		}
	}

	paginator.SetDefault()
	fn := func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(paginator.Offset()).Limit(paginator.PageSize)
	}

	tx = tx.Order(clause.OrderByColumn{Column: clause.Column{Name: paginator.Column}, Desc: paginator.Desc}).Scopes(fn)
	if err = tx.Find(dest).Error; err != nil {
		return err
	}
	return nil
}
