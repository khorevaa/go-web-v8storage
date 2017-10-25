package datamodels

import "github.com/jinzhu/gorm"

type Crserver struct {
	gorm.Model
	Address   string
	Version   string
	BinPath   string
	Port      string
	Range     string
	Directory string
	Tags      []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	gorm.Model
	tag       string
	Crservers []*Crserver `orm:"reverse(many)"`
}
