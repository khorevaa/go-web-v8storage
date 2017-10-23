package datamodels

import "github.com/jinzhu/gorm"

type project struct {
	gorm.Model
	Key         string `json:"key" form:"key"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Picture     []byte
	ShortKey    string
}

type defaultStorageUsers struct {
	gorm.Model
	Project      *project
	Login        string
	Password     string
	Role         string
	User         *User
	BranchRegexp string
}
