package datasource

import (
	//"fmt"

	"../bootstrap"
	"../datamodels"
	"github.com/astaxie/beego"
	"github.com/go-pg/pg"
	_ "github.com/go-pg/pg/orm"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	Init()

}

var DB *pg.DB

func Init() {

	DB = getDBConnection()

	RegisterModels(new(datamodels.Project),
		new(datamodels.User),
		new(datamodels.Crserver),
		new(datamodels.Storage),
		new(datamodels.StorageUser),
		new(datamodels.StorageHistory),
		new(datamodels.StorageInfo),
		new(datamodels.Tag),
		new(datamodels.DefaultStorageUsers))

	//DB.SetMaxIdleConns(10)
	//DB.DB().SetMaxOpenConns(100)

	AutoMigrate()

	initFakeData()
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

func getDBConnection() *pg.DB {

	dbHost := beego.AppConfig.String("db.host")
	dbUser := beego.AppConfig.String("db.user")
	dbName := beego.AppConfig.String("db.name")
	dbPass := beego.AppConfig.String("db.password")
	conn := pg.Connect(&pg.Options{
		Addr:     dbHost,
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
	})

	//if err != nil {
	//	panic(err.Error())
	//}

	return conn
}

//INSERT INTO newTable
//SELECT * FROM oldTable
