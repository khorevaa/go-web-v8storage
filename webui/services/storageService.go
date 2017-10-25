package services

import "../models"
import (
	"../datasource"
	"github.com/jinzhu/gorm"
)

type StorageService interface {
	GetAll() []models.Storage
	GetByID(id int64) (models.Storage,)
	DeleteByID() []models.Storage
	UpdateByID() ([]models.Storage)
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

	s.db.Find(r)
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
