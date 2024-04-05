package models

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ImagePullSecret k8s image pull secret
type ImagePullSecret struct {
	gorm.Model
	// Display name
	Name string `gorm:"type:varchar(60);uniqueIndex:idx_name,priority:1;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	// K8S secret resource name
	SecretName string `gorm:"type:varchar(100);uniqueIndex:idx_name_namespace,priority:1;not null" binding:"required,hostname_rfc1123,min=3,max=100" form:"secretName" json:"secretName,omitempty"`
	// K8S namespace
	SecretNamespace string `gorm:"type:varchar(60);uniqueIndex:idx_name_namespace,priority:2;not null" binding:"required,hostname_rfc1123,min=3,max=60" form:"secretNamespace" json:"secretNamespace,omitempty"`
	// Server location for Docker registry, default https://index.docker.io/v1/
	DockerServer string `gorm:"type:varchar(100)" binding:"required,url,max=100" form:"dockerServer" json:"dockerServer,omitempty"`
	// Username for Docker registry authentication
	DockerUsername string `gorm:"type:varchar(50)" binding:"required,max=50" form:"dockerUsername" json:"dockerUsername,omitempty"`
	// Password for Docker registry authentication
	DockerPassword string `gorm:"type:varchar(100)" binding:"required,max=100" form:"dockerPassword" json:"dockerPassword,omitempty"`
	// Email for Docker registry
	DockerEmail string `gorm:"type:varchar(100)" binding:"email,max=100" form:"dockerEmail" json:"dockerEmail,omitempty"`
	DelAt       int64  `gorm:"uniqueIndex:idx_name_namespace,priority:3;uniqueIndex:idx_name,priority:2;not null" json:"delAt,omitempty"`
}

func (i *ImagePullSecret) BeforeDelete(tx *gorm.DB) (err error) {
	i.DelAt = time.Now().UnixMicro()
	if err = tx.Updates(i).Error; err != nil {
		return
	}
	return
}

func (i *ImagePullSecret) Create(fn func() error) (created *ImagePullSecret, err error) {
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

	if err = tx.Create(i).Error; err != nil {
		return
	}

	if err = fn(); err != nil {
		return
	}

	if err = tx.Commit().Error; err != nil {
		return
	}

	created = i
	return
}

func (i *ImagePullSecret) Update(fn func() error) (updated *ImagePullSecret, err error) {
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

	if err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&ImagePullSecret{}, i.ID).Error; err != nil {
		return
	}

	if err = tx.Updates(i).Error; err != nil {
		return
	}

	if err = fn(); err != nil {
		return
	}

	if err = tx.Commit().Error; err != nil {
		return
	}

	updated = i
	return
}

func (i *ImagePullSecret) Detail() (detail *ImagePullSecret, err error) {
	if err = db.First(i, i.ID).Error; err != nil {
		return
	}

	detail = i
	return
}
