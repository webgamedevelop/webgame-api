package models

import "gorm.io/gorm"

// IngressClass k8s ingress class
type IngressClass struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);uniqueIndex;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	ClassName string `gorm:"type:varchar(20);not null" binding:"required,max=20" form:"className" json:"className,omitempty"`
}

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

// ResourceSpec k8s resource spec
type ResourceSpec struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);uniqueIndex;not null" binding:"required,max=20" form:"name" json:"name,omitempty"`
	Cpu    string `gorm:"type:varchar(12);not null" binding:"required,max=12,k8sCpu" form:"cpu" json:"cpu,omitempty"`
	Memory string `gorm:"type:varchar(12);not null" binding:"required,max=12,k8sMemory" form:"memory" json:"memory,omitempty"`
}

// Repository image repository
type Repository struct {
	gorm.Model
	Name     string `gorm:"type:varchar(50);uniqueIndex;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	RepoName string `gorm:"type:varchar(120);uniqueIndex;not null" binding:"required,max=120" form:"repoName" json:"repoName,omitempty"`
	Tags     []Tag  `json:"tags"`
}

// Tag image tag
type Tag struct {
	gorm.Model
	TagName      string `gorm:"type:varchar(100);uniqueIndex;not null" binding:"required,max=100,semver" form:"tagName" json:"tagName,omitempty"`
	Comment      string `gorm:"type:varchar(200);not null" binding:"max=200" form:"comment" json:"comment,omitempty"`
	RepositoryID uint   `binding:"required" form:"repositoryID" json:"repositoryID"`
}
