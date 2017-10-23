package controllers

import (
	"./api"
	"github.com/labstack/echo"
)

type Контроллер struct {
	HTTPМетод  string
	Путь       string
	Группа     string
	Обработчик echo.HandlerFunc
}

var СписокКонтроллеров []Контроллер

func ЗарегистрироватьКонтроллеры(e *echo.Echo) {

	e.Group("/api/", api.Middleware)

}
