package datamodels

type Project struct {
	BaseModel
	Key         string `json:"key" form:"key"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Picture     []byte
	ShortKey    string
	Tags        []*Tag `pg:",polymorphic:Trackable"`
}

type DefaultStorageUsers struct {
	BaseModel
	Project      Project
	ProjectID    int64
	Login        string
	Password     string
	Role         string
	User         User
	UserID       int64
	BranchRegexp string
}
