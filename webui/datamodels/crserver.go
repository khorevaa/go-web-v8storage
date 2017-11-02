package datamodels

type Crserver struct {
	BaseModel
	Address   string
	Version   string
	BinPath   string
	Port      string
	Range     string
	Directory string
	Tags      []string
}

type Tag struct {
	TrackableID   int64	  `sql:",pk"`
	TrackableType string  `sql:",pk"`
	Text          string
}
