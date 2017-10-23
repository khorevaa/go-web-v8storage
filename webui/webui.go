package main

import (
	"./bootstrap"
	"./middleware/identity"
	"./routes"
)

var app = bootstrap.New("Awesome App", "kataras2006@hotmail.com",
	identity.Configure,
	routes.Configure,
)

func init() {
	app.Bootstrap()
}

func main() {
	app.Listen(":8080")
}