package controllers

import (
	"../services"
	"../models"

)

// StorageController is our /movies controller.
type StorageController struct {
	// mvc.C is just a lightweight lightweight alternative
	// to the "mvc.Controller" controller type,
	// use it when you don't need mvc.Controller's fields
	// (you don't need those fields when you return values from the method functions).
	BaseController

	// Our MovieService, it's an interface which
	// is binded from the main application.
	Service services.StorageService
}


// Get returns list of the movies.
// Demo:
// curl -i http://localhost:8080/movies
//
// The correct way if you have sensitive data:
// func (c *StorageController) Get() (results []viewmodels.Movie) {
// 	data := c.Service.GetAll()
//
// 	for _, movie := range data {
// 		results = append(results, viewmodels.Movie{movie})
// 	}
// 	return
// }
// otherwise just return the datamodels.
func (c *StorageController) Get() {
	// return c.Service.GetAll()
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
