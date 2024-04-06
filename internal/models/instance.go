package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	corev1 "k8s.io/api/core/v1"
)

type Instance struct {
	gorm.Model
	// instance display name
	Name string `gorm:"type:varchar(60);uniqueIndex;not null" binding:"required,max=60" form:"name" json:"name,omitempty"`
	// k8s api resource metadata.name
	ResourceName string `gorm:"type:varchar(100);uniqueIndex:idx_name_namespace,priority:1;not null" binding:"required,hostname_rfc1123,min=3,max=100" form:"resourceName" json:"resourceName,omitempty"`
	// k8s namespace
	ResourceNamespace string `gorm:"type:varchar(60);uniqueIndex:idx_name_namespace,priority:2;not null" binding:"required,hostname_rfc1123,min=3,max=60" form:"resourceNamespace" json:"resourceNamespace,omitempty"`
	// k8s api resource metadata.uid
	ResourceUID string `gorm:"type:varchar(50)" json:"resourceUID,omitempty"`
	// k8s api resource metadata.resourceVersion
	ResourceVersion   string            `gorm:"type:varchar(50)" json:"resourceVersion,omitempty"`
	Labels            datatypes.JSONMap `json:"labels,omitempty"`
	Annotations       datatypes.JSONMap `json:"annotations,omitempty"`
	GameTypeVersionID uint              `binding:"required" form:"gameTypeVersionID" json:"gameTypeVersionID"`
	GameTypeVersion   GameTypeVersion
	DomainPrefix      string `gorm:"type:varchar(50);uniqueIndex" binding:"hostname_rfc1123,max=50" form:"domainPrefix" json:"domainPrefix"`
	IndexPage         string `gorm:"type:varchar(50);uniqueIndex" binding:"max=50" form:"indexPage" json:"indexPage"`
	IngressClassID    uint   `binding:"required" form:"ingressClassID" json:"ingressClassID"`
	IngressClass      IngressClass
	Replicas          int32  `binding:"required" form:"replicas" json:"replicas"`
	Image             string `form:"image" json:"image"`
	ResourceSpecID    uint   `binding:"required" form:"resourceSpecID" json:"resourceSpecID,omitempty"`
	ResourceSpec      ResourceSpec
	ImagePullSecrets  datatypes.JSONSlice[corev1.LocalObjectReference] `form:"imagePullSecrets" json:"imagePullSecrets"`
	Envs              datatypes.JSONSlice[corev1.EnvVar]               `form:"envs" json:"envs"`
	Synced            bool                                             `form:"synced" json:"synced,omitempty"`
	Version           optimisticlock.Version
}
