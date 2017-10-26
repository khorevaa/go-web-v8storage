package datamodels

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Storage struct {
	gorm.Model
	Project          *Project
	ProjectID 		 int64
	Crserver         *Crserver
	CrserverID 		 int64
	BranchName       string
	Description      string       `json:"Description" form:"Description"`
	User          	 *User
	UserID 		 	 int64
	HistoryUpdatedAt time.Time
	ConnectLogin     string
	ConnectPassword  string
	Disable          bool
	Type             string
	StorageInfo      StorageInfo
	StorageUsers     []*StorageUser
	StorageHistory   []*StorageHistory
}

type StorageInfo struct {
	gorm.Model
	StorageID      int64
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
