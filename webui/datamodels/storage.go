package datamodels

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Storage struct {
	gorm.Model
	Project          *Project
	Crserver         *Crserver
	BranchName       string
	Description      string       `json:"Description" form:"Description"`
	User_ID          *User
	HistoryUpdatedAt time.Time
	ConnectLogin     string
	ConnectPassword  string
	Disable          bool
	Type             string
	StorageInfo      *StorageInfo `orm:"null;rel(one);on_delete(set_null)"`
	StorageUsers     []*StorageUser
	StorageHistory   []*StorageHistory
}

type StorageInfo struct {
	ID             int64
	Storage        *Storage `orm:"reverse(one)"`
	ProductName    string
	BasicRelease   string
	CurrentRelease string
	BasicCommit    string
}

type StorageHistory struct {
	Storage    *Storage
	VersionNum int64
	DateTime   time.Time
	Author     *StorageUser
	Comment    string
}

type StorageUser struct {
	gorm.Model
	Storage *Storage
	Login   string
	Role    string
	Disable bool
	User    *User
	Alias   string
}
