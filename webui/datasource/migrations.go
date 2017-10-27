package datasource

import "github.com/go-pg/pg/orm"

var models []interface{}
var migrations []string

func RegisterModel(model interface{}) {
	models = append(models, model)
}

func RegisterModels(m ...interface{}) {
	models = m
}

func RegisterMigration(migration string) {
	migrations = append(migrations, migration)
}

func AutoMigrate() {
	for _, migration := range migrations {
		DB.Exec(migration)
	}

	for _, model := range models{
		err := DB.CreateTable(model, &orm.CreateTableOptions{
			Temp: false,
			IfNotExists: true,

		})
		if err != nil {
			panic(err)
		}
	}

	//DB.AutoMigrate(models...)
}
