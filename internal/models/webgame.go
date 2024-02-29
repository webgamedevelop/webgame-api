package models

import (
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"

	corev1 "k8s.io/api/core/v1"
)

type Webgame struct {
	gorm.Model
	DisplayName      string
	GameType         string
	Domain           string
	IndexPage        string
	IngressClass     string
	ServerPort       int32
	Replicas         int32
	Image            string
	ImagePullSecrets datatypes.JSONSlice[corev1.LocalObjectReference]
	Version          optimisticlock.Version
}

type MutateFn func(interface{}) error

// MutateWebgame is sample for MutateFn
func MutateWebgame(dest interface{}) error {
	webgame, ok := dest.(*Webgame)
	if !ok {
		return fmt.Errorf("type assertion failure for %v", dest)
	}

	webgame.Replicas += 1
	return nil
}

func Update(dest interface{}, mutate MutateFn) error {
	for {
		if err := db.First(dest).Error; err != nil {
			return err
		}
		if err := mutate(dest); err != nil {
			return err
		}
		tx := db.Updates(dest)
		if err := tx.Error; err != nil {
			return err
		}
		if tx.RowsAffected != 0 {
			break
		}
	}
	return nil
}
