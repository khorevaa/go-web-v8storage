package services

import (
	"../datasource"
	//"../models"
	models "../datamodels"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type ProjectService interface {
	GetAll() []models.Project
	GetByID(id int64) models.Project
	DeleteByID() []models.Project
	UpdateByID() []models.Project
	GetList(page int, pageSize int) (r []models.Project, c int)
}

// NewMovieService returns the default movie service.
func NewProjectService() *projectService{
	return &projectService{
		db:        datasource.DB,
		tableName: datasource.TableName("Storage"),
	}
}

type projectService struct {
	db        *pg.DB
	tableName string
}

func (s *projectService) GetAll() (projects []models.Project) {

	_, err := s.db.Model(&projects).Order("id").SelectAndCount(&projects)

	if err != nil {
		panic(err)
	}
	return
}

func (s *projectService) GetList(page int, pageSize int) (storages []models.Project, count int) {

	offset := (page - 1) * pageSize

	//f := []datamodels.Storage{}

	//rows, err := s.db.Raw("SELECT * FROM storages").Rows()
	//
	//if err != nil  {
	//	panic(err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	////rows.Scan(&name, &age, &email)
	//}

	//s.db.Last(&f)

	count, err := orm.NewQuery(s.db, &storages).Order("id").Limit(pageSize).Offset(offset).Column("storage.*", "Crserver", "Project").SelectAndCount(&storages)

	//count, err:= s.db.Model(&storages).Order("id").Limit(pageSize).Offset(offset).Column("storage.*","Tags", "Crserver", "Crserver.Tags", "Project", "Project.Tags").SelectAndCount(&storages)
	//println(storages[2].Crserver.Tags[0].Text)
	if err != nil {
		panic(err)
	}

	return
}

func (s *projectService) GetByID(id int64) (storage models.Project) {

	err:= s.db.Model(&storage).Order("id").Column("storage.*").Where("id =?", id).Select(&storage)

	if err != nil {
		panic(err)
	}
	return
}

func (s *projectService) Update(storage models.Project) (err error) {

	_, err= s.db.Model(&storage).Update(&storage)

	return
}

func (s *projectService) DeleteByID() (r []models.Project) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}

func (s *projectService) UpdateByID() (r []models.Project) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}
