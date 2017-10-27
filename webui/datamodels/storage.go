package datamodels

import (
	"time"
)

type Storage struct {
	BaseModel
	Project          *Project
	ProjectID        int64
	Crserver         *Crserver
	CrserverID       int64
	BranchName       string
	Description      string `json:"Description" form:"Description"`
	User             *User
	UserID           int64
	HistoryUpdatedAt time.Time
	ConnectLogin     string
	ConnectPassword  string
	Disable          bool
	Type             string
	StorageInfo      StorageInfo
	StorageInfoID    int64
	StorageUsers     []*StorageUser
	StorageHistory   []*StorageHistory
	Tags             []*Tag `pg:",polymorphic:Trackable"`
}

type StorageInfo struct {
	BaseModel
	Storage        *Storage
	StorageID      int64
	ProductName    string
	BasicRelease   string
	CurrentRelease string
	BasicCommit    string
}

type StorageHistory struct {
	Storage    *Storage
	StorageID  int64
	VersionNum int64
	DateTime   time.Time
	Author     *StorageUser
	AuthorID   int64
	Comment    string
}

type StorageUser struct {
	BaseModel
	Storage   *Storage
	StorageID int64
	Login     string
	Role      string
	Disable   bool
	User      *User
	UserID    int64
	Alias     string
}
