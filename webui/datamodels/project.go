package datamodels

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Key         string `json:"key" form:"key"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Picture     []byte
	ShortKey    string
}

type DefaultStorageUsers struct {
	gorm.Model
	Project      *Project
	Login        string
	Password     string
	Role         string
	User         *User
	BranchRegexp string
}
