package httpserver

import (
	"betsys/config"
	"betsys/service/serverservice"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config        config.Config
	serverService serverservice.Service
}

func New(config config.Config, serverService serverservice.Service) Server {
	return Server{
		config:        config,
		serverService: serverService,
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

	serverGroup := e.Group("/server")
	serverGroup.POST("/register", s.ServerRegister)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)))
}
