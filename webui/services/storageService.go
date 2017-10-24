package services

import "../models"
import (
	"../datasource"
	"github.com/jinzhu/gorm"
)

type StorageService interface {
	GetAll() []models.Storage
	GetByID(id int64) (models.Storage, bool)
	DeleteByID(id int64) bool
	UpdateByID(id int64, poster string, genre string) (models.Storage, error)
}

// NewMovieService returns the default movie service.
func NewStorageService() *storageService {
	return &storageService{
		db: datasource.DB,
	}
}

type storageService struct {
	db *gorm.DB
}

func (s *storageService) GetAll() (r []models.Storage) {

	s.db.Find(&r)

	return
}
