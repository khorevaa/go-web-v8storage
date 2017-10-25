package datasource

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"../datamodels"
	"../bootstrap"
	"github.com/jinzhu/gorm"
	"fmt"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	Init()

}

var DB *gorm.DB

func Connect(host string, database string, user string, pass string) (db *gorm.DB, err error) {
	dbConnString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, database, pass)
	db, err = gorm.Open("postgres", dbConnString)
	return
}

func Init() {

	DB = getDBConnection()

	RegisterModels(new(datamodels.Project),
		new(datamodels.User),
		new(datamodels.Crserver),
		new(datamodels.Storage),
		new(datamodels.StorageUser),
		new(datamodels.StorageHistory),
		new(datamodels.StorageInfo),
		new(datamodels.DefaultStorageUsers))

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	AutoMigrate()
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func getDBConnection() *gorm.DB {

	dbHost := beego.AppConfig.String("db.host")
	dbUser := beego.AppConfig.String("db.user")
	dbName := beego.AppConfig.String("db.name")
	dbPass := beego.AppConfig.String("db.password")
	conn, err := Connect(dbHost, dbName, dbUser, dbPass)
	if err != nil {
		panic(err.Error())
	}

	return conn
}