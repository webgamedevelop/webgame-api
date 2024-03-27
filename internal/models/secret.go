package models

import "gorm.io/gorm"

// ImagePullSecret k8s image pull secret
type ImagePullSecret struct {
	gorm.Model
	Name            string `gorm:"type:varchar(60);uniqueIndex;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	SecretName      string `gorm:"type:varchar(100);uniqueIndex:idx_name_namespace,priority:1;not null" binding:"required,hostname_rfc1123,min=3,max=100" form:"secretName" json:"secretName,omitempty"`
	SecretNamespace string `gorm:"type:varchar(60);uniqueIndex:idx_name_namespace,priority:2;not null" binding:"required,hostname_rfc1123,min=3,max=60" form:"secretNamespace" json:"secretNamespace,omitempty"`
	DockerServer    string `gorm:"type:varchar(100)" binding:"required,hostname_rfc1123,max=100" form:"dockerServer" json:"dockerServer,omitempty"`
	DockerUsername  string `gorm:"type:varchar(50)" binding:"required,max=50" form:"dockerUsername" json:"dockerUsername,omitempty"`
	DockerPassword  string `gorm:"type:varchar(100)" binding:"required,max=100" form:"dockerPassword" json:"dockerPassword,omitempty"`
	DockerEmail     string `gorm:"type:varchar(100)" binding:"email,max=100" form:"dockerEmail" json:"dockerEmail,omitempty"`
}
