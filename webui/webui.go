package webui

import (
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server implements an HTTP server which exposes a JSON API to view and manage gocraft/work items.
type Server struct {
	namespace string
	hostPort  string
	server    *echo.Echo
}

type Context struct {
	echo.Context
	*Server
	Session map[string]string
}

// NewServer creates and returns a new server. The 'namespace' param is the redis namespace to use. The hostPort param is the address to bind on to expose the API.
func NewServer(namespace string, pool *redis.Pool, hostPort string) *Server {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := &Server{
		namespace: namespace,
		hostPort:  hostPort,
		server:    e,
	}

	// Debug mode
	e.Debug = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}, AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE}}))

	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	return server
}

// Start starts the server listening for requests on the hostPort specified in NewServer.
func (w *Server) Start() {
	w.server.Logger.Fatal(w.server.Start(w.hostPort))
}

// Stop stops the server and blocks until it has finished.
func (w *Server) Stop() {
	w.server.Logger.Fatal(w.server.Close())
}
