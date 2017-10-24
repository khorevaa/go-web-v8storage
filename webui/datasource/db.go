package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	c "../config"
)

var DB *gorm.DB

func init() {
	dbinfo := "host=" + c.Config.Postgres.Host +
		" dbname=" + c.Config.Postgres.Database +
		" user=" + c.Config.Postgres.User +
		" password=" + c.Config.Postgres.Password +
		" sslmode="

	if c.Config.Postgres.SSL {
		dbinfo += "enable"
	} else {
		dbinfo += "disable"
	}

	db, err := gorm.Open("postgres", dbinfo)

	if err != nil {
		panic("failed to connect to database")
	}

	DB = db

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
