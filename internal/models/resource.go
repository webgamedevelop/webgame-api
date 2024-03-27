package models

import "gorm.io/gorm"

// IngressClass k8s ingress class
type IngressClass struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);uniqueIndex;not null" binding:"required,max=50" form:"name" json:"name,omitempty"`
	ClassName string `gorm:"type:varchar(20);not null" binding:"required,max=20" form:"className" json:"className,omitempty"`
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
