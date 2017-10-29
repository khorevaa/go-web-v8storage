package datasource

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
	for _, model := range models {
		AutoMigrateModel(model)
	}

	//DB.AutoMigrate(models...)
}
