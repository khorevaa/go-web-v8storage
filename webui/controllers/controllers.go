package controllers

import (
	"github.com/labstack/echo"
	api "../controllers/api"
)


type Контроллер struct {
	HTTPМетод string
	Путь string
	Группа string
	Обработчик echo.HandlerFunc
}

var СписокКонтроллеров []Контроллер

func ЗарегистрироватьКонтроллеры(e *echo.Echo) {


	e.Group("/api/", api.Middleware)


}
