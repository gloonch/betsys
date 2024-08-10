package httpserver

import (
	"betsys/service/serverservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) ServerRegister(c echo.Context) error {

	var req serverservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.serverService.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
