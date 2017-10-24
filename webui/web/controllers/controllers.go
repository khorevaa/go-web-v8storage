package controllers

import (
	"../../bootstrap"
	//"../middleware"
	"../../services"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {

	b.Controller("/storages",
		new(StorageController),
		services.NewStorageService())

}