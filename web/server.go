package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manattan/npb_api/handler"
	"github.com/manattan/npb_api/usecase"
)

func NewServer(teamUC *usecase.TeamUseCase) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	teamHandler := handler.NewTeamHandler(teamUC)

	v0 := e.Group("/api/v0")

	v0.GET("/team/:id", teamHandler.GetTeam)

	return e
}
