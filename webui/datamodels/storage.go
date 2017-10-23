package datamodels

import (
	"github.com/jinzhu/gorm"
	"time"
)

type storage struct {
	gorm.Model
	ID               int64  `json:"id" form:"id"`
	Project          *project
	Crserver         *crserver
	BranchName       string
	Description      string `json:"Description" form:"Description"`
	User_ID          *User
	HistoryUpdatedAt time.Time
	ConnectLogin     string
	ConnectPassword  string
	Disable          bool
	Type             string
	StorageInfo      *storageInfo
	StorageUsers     []storageUser
	StorageHistory   []storageHistory
}

type storageInfo struct {
	ID             int64
	Storage        storage
	ProductName    string
	BasicRelease   string
	CurrentRelease string
	BasicCommit    string
}

type storageHistory struct {
	Storage    storage
	VersionNum int64
	DateTime   time.Time
	Author     storageUser
	Comment    string
}

type storageUser struct {
	gorm.Model
	ID      int64 `json:"id" form:"id"`
	Storage storage
	Login   string
	Role    string
	Disable bool
	User    *User
	Alias   string
}
