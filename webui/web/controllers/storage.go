package controllers

import (
	"github.com/kataras/iris/mvc"
	"../../services"
	"github.com/kataras/iris"
	datamodels "../../models"
)

// StorageController is our /movies controller.
type StorageController struct {
	// mvc.C is just a lightweight lightweight alternative
	// to the "mvc.Controller" controller type,
	// use it when you don't need mvc.Controller's fields
	// (you don't need those fields when you return values from the method functions).
	mvc.C

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
func (c *StorageController) Get() (mvc.Result) {
	return mvc.View{
		Name: "wviews/index.html",
		Data: "",
	}
	// return c.Service.GetAll()
}

// GetBy returns a movie.
// Demo:
// curl -i http://localhost:8080/movies/1
func (c *StorageController) GetBy(id int64) (movie datamodels.Storage, found bool) {
	return c.Service.GetByID(id) // it will throw 404 if not found.
}

// DeleteBy deletes a movie.
// Demo:
// curl -i -X DELETE -u admin:password http://localhost:8080/movies/1
func (c *StorageController) DeleteBy(id int64) interface{} {
	wasDel := c.Service.DeleteByID(id)
	if wasDel {
		// return the deleted movie's ID
		return iris.Map{"deleted": id}
	}
	// right here we can see that a method function can return any of those two types(map or int),
	// we don't have to specify the return type to a specific type.
	return iris.StatusBadRequest
}
