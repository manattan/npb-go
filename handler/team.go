package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/npb_api/usecase"
)

type TeamHandler struct {
	teamUC *usecase.TeamUseCase
}

func NewTeamHandler(teamUC *usecase.TeamUseCase) *TeamHandler {
	return &TeamHandler{teamUC: teamUC}
}

func (h *TeamHandler) GetTeams(c echo.Context) error {
	teams, err := h.teamUC.GetTeams()

	if err != nil {
		return c.JSON(http.StatusBadRequest, &getTeamError{
			Message: "hoge error",
		})
	}

	return c.JSON(http.StatusOK, &teams)
}

func (h *TeamHandler) GetTeam(c echo.Context) error {

	teamId := c.Param("id")

	team, err := h.teamUC.GetTeam(teamId)

	if err != nil {
		return c.JSON(http.StatusNotFound, &getTeamError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &getTeamRes{
		ID:       team.ID,
		Name:     team.Name,
		Initial:  team.Initial,
		LeagueId: team.LeagueId,
	})

}

type getTeamError struct {
	Message string `json:"message"`
}

type getTeamRes struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initial  string `json:"initial"`
	LeagueId int    `json:"leagueId"`
}
