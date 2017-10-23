package datamodels

import "github.com/jinzhu/gorm"

type crserver struct {
	gorm.Model
	ID        int64
	Address   string
	Version   string
	BinPath   string
	Port      string
	Range     string
	Directory string
	Tags      []string
}
