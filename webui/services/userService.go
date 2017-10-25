package services

import models "../datamodels"
import (
	"../datasource"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int64) (models.User)
}

// NewMovieService returns the default movie service.
func NewUserService() userService {
	return userService{
		db:        datasource.DB,
		tableName: datasource.TableName("User"),
	}
}

type userService struct {
	db        *gorm.DB
	tableName string
}

func (s *userService) GetAll() (r []models.User) {

	s.db.Find(r)

	return
}

func (s *userService) GetByID(id int64) (r models.User) {

	//s.orm.QueryTable(s.tableName).OrderBy("-id").Filter("id", id).One(r)

	return
}
