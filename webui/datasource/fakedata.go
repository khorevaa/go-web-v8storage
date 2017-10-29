package datasource

import (
	"fmt"

	"time"

	"reflect"

	"strings"

	"../datamodels"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var dropTableOpt = &orm.DropTableOptions{
	true,
	false,
}

var createTableOpt = &orm.CreateTableOptions{
	IfNotExists: true,
}

func initFakeData() {

	AutoMigrateModel(datamodels.Project{})

	//initProjects()
	//initUsers()
	//initCrservers()

}

func CopyTable(ExistsTableName string, NewTableName string) {

	println(orm.Tables.Get(reflect.TypeOf(datamodels.Project{})).Name)
}

func AutoMigrateModel(model interface{}) {

	curTableName := strings.Replace(string(orm.Tables.Get(reflect.TypeOf(datamodels.Project{})).Name), "\"", "", 2)
	tempTableName := curTableName + "_temp"

	err := DB.RunInTransaction(func(tx *pg.Tx) error {

		_, errDrop := tx.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s ", tempTableName))
		if errDrop != nil {
			panic(errDrop)
		}

		_, err := tx.Exec(fmt.Sprintf("CREATE TABLE %s AS (SELECT * FROM %s) ", tempTableName, curTableName))

		if err != nil {
			panic(err)
		}

		tx.DropTable(model, dropTableOpt)
		tx.CreateTable(model, createTableOpt)

		_, err = tx.Exec(fmt.Sprintf("INSERT INTO  %s (SELECT * FROM %s) ", tempTableName, curTableName))

		if err != nil {
			panic(err)
		}
		_, err = tx.Exec(fmt.Sprintf("DROP TABLE %s", tempTableName))

		if err != nil {
			panic(err)
		}
		return err
	})

	if err != nil {
		panic(err)
	}
}

func initProjects() {

	p1 := datamodels.Project{
		BaseModel:   datamodels.BaseModel{ID: 1},
		Key:         "jira",
		Name:        "JIRA Проект",
		Description: "Описание JIRA",
		ShortKey:    "jira",
		Tags:        []string{"dev", "home"},
	}

	p2 := datamodels.Project{
		BaseModel:   datamodels.BaseModel{ID: 2},
		Key:         "bitbucket",
		Name:        "Bitbucket Проект",
		Description: "Описание bitbucket",
		ShortKey:    "git",
		Tags:        []string{"prod", "work"},
	}

	p3 := datamodels.Project{
		BaseModel:   datamodels.BaseModel{ID: 3},
		Key:         "bamboo",
		Name:        "Bamboo Проект",
		Description: "Описание bamboo",
		ShortKey:    "bamboo",
		Tags:        []string{"dev", "work"},
	}

	array := []datamodels.Project{
		p1,
		p2,
		p3,
	}

	fmt.Sprintln(array)

	//r, err := orm.NewQuery(DB, &array).OnConflict("DO NOTHING").Returning("*").Insert()

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}

}

func initUsers() {

	user1 := datamodels.User{
		1,
		"Тестовый пользователь",
		"test_user",
		nil,
		time.Now(),
		false,
	}

	user2 := datamodels.User{
		2,
		"Тестовый admin пользователь",
		"admin",
		nil,
		time.Now(),
		false,
	}

	user3 := datamodels.User{
		3,
		"Тестовый хранилища пользователь",
		"work",
		nil,
		time.Now(),
		false,
	}
	array := []datamodels.User{
		user1,
		user2,
		user3,
	}

	//fmt.Sprintln(array)

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}
}

func initCrservers() {

	crserver1 := datamodels.Crserver{
		BaseModel: datamodels.BaseModel{ID: 1},
		Address:   "dev.storage.1c.can",
		Version:   "8.3.10.2086",
		BinPath:   "/path/to/bin",
		Port:      "1452",
		Range:     "1452-1456",
		Tags:      []string{"dev", "main"},
	}

	crserver2 := datamodels.Crserver{
		BaseModel: datamodels.BaseModel{ID: 2},
		Address:   "release.storage.1c.can",
		Version:   "8.3.8.1806",
		BinPath:   "/path/to/bin",
		Port:      "1452",
		Range:     "1452-1456",
		Tags:      []string{"release"},
	}

	crserver3 := datamodels.Crserver{
		BaseModel: datamodels.BaseModel{ID: 3},
		Address:   "dev2.storage.1c.can",
		Version:   "8.3.11.2306",
		BinPath:   "/path/to/bin",
		Port:      "1452",
		Range:     "1452-1456",
		Tags:      []string{"dev", "additian", "new_platform"},
	}

	array := []datamodels.Crserver{
		crserver1,
		crserver2,
		crserver3,
	}

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}
}
