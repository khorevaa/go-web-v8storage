package main

import (
	"./bootstrap"
	"./middleware/identity"
	"./web/controllers"
)

var app = bootstrap.New("Управление серверами хранилищ 1С ", "horevaa@yandex.ru",
	identity.Configure,
	controllers.Configure,
)

func init() {
	app.Bootstrap()
}

func main() {
	app.Listen(":8080")
}