package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	dbinfo := "host=" + Config.Postgres.Host +
		" dbname=" + Config.Postgres.Database +
		" user=" + Config.Postgres.User +
		" password=" + Config.Postgres.Password +
		" sslmode="

	if Config.Postgres.SSL {
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
