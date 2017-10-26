package services

import "../models"
import (
	"../datasource"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	GetAll() []models.User
	GetByID(id int64) models.User
	UserGetByName(i string) (models.User, error)
}

// NewMovieService returns the default movie service.
func NewUserService() *userService {
	return &userService{
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

func (s *userService) UserGetByName(i string) (u models.User, err error) {

	u = models.User{}

	u.ID = 1
	u.Username = "admin"
	//s.db.Find(u).Where("Username=&u", i)
	return
}

func (s *userService) GetByID(id int64) (r models.User) {

	r = models.User{}

	r.ID = 1
	r.Username = "admin"
	//s.orm.QueryTable(s.tableName).OrderBy("-id").Filter("id", id).One(r)

	return
}
