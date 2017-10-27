package controllers

import (
	"../services"
	models "../datamodels"
	//"../models"
	"../libs"
	"github.com/astaxie/beego"

)

// StorageController is our /movies controller.
type StorageController struct {
	BaseController
	Service services.StorageService
}

func NewStorageController() *StorageController {
	return &StorageController{Service: services.NewStorageService()}
}

func (c *StorageController) Get() {
	// return c.Service.GetAll()
}

func (this *StorageController) List() {
	page, _ := this.GetInt("page")
	if page < 1 {
		page = 1
	}

	list, count := services.NewStorageService().GetList(page, this.pageSize)

	this.Data["pageTitle"] = "Список хранилищ 1С "
	this.Data["list"] = list
	this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("GroupController.List"), true).ToString()
	this.display()
}


// GetBy returns a movie.
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *StorageController) GetBy(id int64) (Storage models.Storage) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}

// DeleteBy deletes a movie.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *StorageController) DeleteBy(id int64) (r interface{}) {

	return
}
