package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// IngressClass k8s ingress class
type IngressClass struct {
	gorm.Model
	// Display name
	Name string `gorm:"type:varchar(50);uniqueIndex:idx_name,priority:1;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	// K8S ingress class resource name
	ClassName string `gorm:"type:varchar(20);not null" binding:"required,max=20" form:"className" json:"className,omitempty"`
	// Imported flag
	Imported bool  `form:"imported" json:"imported,omitempty"`
	DelAt    int64 `gorm:"uniqueIndex:idx_name,priority:2;not null" json:"delAt,omitempty"`
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
