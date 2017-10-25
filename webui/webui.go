package main

import (
	"./bootstrap"
	"./controllers"
	"./datasource"
	//"github.com/astaxie/beego"
)

var app = bootstrap.New(datasource.Configure, controllers.Configure)

func init() {
	app.Bootstrap()
}

func main() {

	app.Listen()
}