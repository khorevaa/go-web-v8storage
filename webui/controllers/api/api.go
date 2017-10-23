package api

import (
	//. "../../../webui"
	"github.com/labstack/echo"
)

type APIMiddleware struct {
	echo.HandlerFunc
}

const prefixAPI  = "/api/"

type Router struct {
	*echo.Router
	routers []string
}

func NewAPIRouter(e *echo.Echo) *echo.Router{

	router := e.Group(prefixAPI, APIMiddleware{}.Handler)

	router.GET("users", )

	return
}

func (a *APIMiddleware) Middleware(config  echo.Context) echo.MiddlewareFunc{

	return a.Handler

}

func (a *APIMiddleware) Handler(next echo.HandlerFunc) echo.HandlerFunc{

	return func(c echo.Context) error {

		return next(c)
		}
	}
}

func Middleware (c echo.Context) (e error) {

	return
}


func