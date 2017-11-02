package datasource

import "github.com/go-pg/pg/orm"

var models []interface{}
var migrations []string

var dropTableOpt = &orm.DropTableOptions{
	true,
	false,
}

var createTableOpt = &orm.CreateTableOptions{
	IfNotExists: true,
}

func RegisterModel(model interface{}) {
	models = append(models, model)
}

func RegisterModels(m ...interface{}) {
	models = m
}

func RegisterMigration(migration string) {
	migrations = append(migrations, migration)
}


func AutoMigrateModel(model interface{}) {

	DB.DropTable(model, dropTableOpt)
	DB.CreateTable(model, createTableOpt)

}
func AutoMigrate() {
	for _, migration := range migrations {
		DB.Exec(migration)
	}
	for _, model := range models {
		AutoMigrateModel(model)
	}

	//DB.AutoMigrate(models...)
}
