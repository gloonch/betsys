package httpserver

import (
	"betsys/config"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	//serverService	serverservice.Service
}

func New(config config.Config) Server {
	return Server{
		config: config,
	}
}

func (s Server) Serve() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Interesting! <-----
	//e.GET("/health-check", func(c echo.Context) error {
	//
	//})

	e.GET("/health-check", s.healthCheck)

	//serverGroup := e.Group("/server")
	//serverGroup.POST("register", s.)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
