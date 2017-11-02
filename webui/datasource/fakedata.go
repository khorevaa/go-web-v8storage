package datasource

import (
	"fmt"
	"time"
	"../datamodels"
)



func initFakeData() {

	AutoMigrate()

	initProjects()
	initUsers()
	initCrservers()
	initStorages()
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

func initStorages() {

	p1 := datamodels.Storage{
		BaseModel:        datamodels.BaseModel{ID: 1},
		ProjectID:        1,
		CrserverID:       1,
		BranchName:       "master",
		Description:      "Рабочение хранилище JIRA",
		UserID:           1,
		HistoryUpdatedAt: time.Now(),
		ConnectLogin:     "Администратор",
		ConnectPassword:  "Пароль",
		Type:             "http",
		StorageInfoID:    1,
		Tags:             []string{"prod", "work"},
	}

	p2 := datamodels.Storage{
		BaseModel:   datamodels.BaseModel{ID: 2},
		ProjectID:        2,
		CrserverID:       1,
		BranchName:       "master",
		Description:      "Рабочение хранилище Bitbucket",
		UserID:           1,
		HistoryUpdatedAt: time.Now(),
		ConnectLogin:     "Администратор",
		ConnectPassword:  "Пароль",
		Type:             "http",
		StorageInfoID:    2,
		Tags:        	 []string{"prod", "work"},
	}

	p3 := datamodels.Storage{
		BaseModel:   datamodels.BaseModel{ID: 3},
		ProjectID:        3,
		CrserverID:       1,
		BranchName:       "master",
		Description:      "Рабочение хранилище Bamboo",
		UserID:           1,
		HistoryUpdatedAt: time.Now(),
		ConnectLogin:     "Администратор",
		ConnectPassword:  "Пароль",
		Type:             "http",
		StorageInfoID:    3,
		Tags:        	 []string{"prod", "work"},
	}


	p4 := datamodels.Storage{
		BaseModel:        datamodels.BaseModel{ID: 4},
		ProjectID:        1,
		CrserverID:       1,
		BranchName:       "feature/new-ss",
		Description:      "Задача по разработки New-SS (JIRA)",
		UserID:           2,
		HistoryUpdatedAt: time.Now(),
		ConnectLogin:     "Администратор",
		ConnectPassword:  "Пароль",
		Type:             "http",
		StorageInfoID:    1,
		Tags:             []string{"dev", "new"},
	}

	array := []datamodels.Storage{
		p1,
		p2,
		p3,
		p4,
	}

	fmt.Sprintln(array)

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}

}

func initDefaultStorageUsers() {

	p1 := datamodels.Project{
	}

	p2 := datamodels.Project{
	}

	p3 := datamodels.Project{
	}

	array := []datamodels.Project{
		p1,
		p2,
		p3,
	}

	fmt.Sprintln(array)

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}

}

func initStorageHistory() {

	p1 := datamodels.Project{
	}

	p2 := datamodels.Project{
	}

	p3 := datamodels.Project{
	}

	array := []datamodels.Project{
		p1,
		p2,
		p3,
	}

	fmt.Sprintln(array)

	_, err := DB.Model(&array).OnConflict("DO NOTHING").Returning("*").Insert()
	if err != nil {
		panic(err)
	}

}
