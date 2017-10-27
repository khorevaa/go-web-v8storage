package services

import (
	"../datasource"
	"../models"
	"github.com/jinzhu/gorm"
	//"../datamodels"
)

type StorageService interface {
	GetAll() []models.Storage
	GetByID(id int64) models.Storage
	DeleteByID() []models.Storage
	UpdateByID() []models.Storage
	GetList(page int, pageSize int) (r []models.Storage, c int64)
}

// NewMovieService returns the default movie service.
func NewStorageService() *storageService {
	return &storageService{
		db:        datasource.DB,
		tableName: datasource.TableName("Storage"),
	}
}

type storageService struct {
	db        *gorm.DB
	tableName string
}

func (s *storageService) GetAll() (r []models.Storage) {

	s.db.Order("ID").Find(&r)
	return
}

func (s *storageService) GetList(page int, pageSize int) (storages []models.Storage, count int64) {

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
	s.db.Preload("Project").Order("id").Limit(pageSize).Offset(offset).Find(&storages).Count(&count)
	return
}

func (s *storageService) GetByID(id int64) (r models.Storage) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").Filter("id", id).One(r)

	return
}

func (s *storageService) DeleteByID() (r []models.Storage) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}

func (s *storageService) UpdateByID() (r []models.Storage) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").All(r)

	return
}
