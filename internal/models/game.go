package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GameType struct {
	gorm.Model
	Name           string       `gorm:"type:varchar(60);uniqueIndex" binding:"required,max=60" form:"name" json:"name,omitempty"`
	TypeName       string       `gorm:"type:varchar(60);uniqueIndex" binding:"required,max=60" form:"typeName" json:"typeName,omitempty"`
	RootDomain     string       `gorm:"type:varchar(200);uniqueIndex" binding:"required,max=200,hostname_rfc1123" form:"rootDomain" json:"rootDomain"`
	ResourceSpecID uint         `binding:"required" form:"resourceSpecID" json:"resourceSpecID"`
	RepositoryID   uint         `binding:"required" form:"repositoryID" json:"repositoryID"`
	ResourceSpec   ResourceSpec `json:"resourceSpec"`
	Repository     Repository   `json:"repository"`
}

type GameTypeVersion struct {
	gorm.Model
	Name           string                      `gorm:"type:varchar(100);uniqueIndex;not null" binding:"required,hostname_rfc1123,min=3,max=100" form:"name" json:"name,omitempty"`
	GameTypeID     uint                        `gorm:"uniqueIndex:idx_type_version,priority:1;not null" binding:"required" form:"gameTypeID" json:"gameTypeID,omitempty"`
	GameType       GameType                    `json:"gameType"`
	GameVersion    string                      `gorm:"type:varchar(100);uniqueIndex:idx_type_version,priority:2;not null" binding:"required,max=100,semver" form:"gameVersion" json:"gameVersion,omitempty"`
	ResourceSpecID uint                        `binding:"required" form:"resourceSpecID" json:"resourceSpecID,omitempty"`
	ResourceSpec   ResourceSpec                `json:"resourceSpec"`
	RepositoryID   uint                        `binding:"required" form:"repositoryID" json:"repositoryID"`
	Repository     Repository                  `json:"repository"`
	ImageTagID     uint                        `binding:"required" form:"ImageTagID" json:"ImageTagID,omitempty"`
	ImageTag       Tag                         `gorm:"foreignKey:ImageTagID"`
	ServerPort     int32                       `binding:"required" form:"serverPort" json:"serverPort,omitempty"`
	Command        datatypes.JSONSlice[string] `form:"command" json:"command,omitempty"`
	Args           datatypes.JSONSlice[string] `form:"args" json:"args,omitempty"`
}
