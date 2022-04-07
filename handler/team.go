package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/npb_go/usecase"
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
		return c.JSON(http.StatusInternalServerError, &GetTeamError{
			Message: err.Error(),
		})
	}

	var res []*GetTeamSuccess
	for _, team := range teams {
		res = append(res, &GetTeamSuccess{
			ID:       team.ID,
			Name:     team.Name,
			Initial:  team.Initial,
			LeagueId: team.LeagueId,
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *TeamHandler) GetTeam(c echo.Context) error {

	teamId := c.Param("id")

	team, err := h.teamUC.GetTeam(teamId)

	if err != nil {
		return c.JSON(http.StatusNotFound, &GetTeamError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &GetTeamSuccess{
		ID:       team.ID,
		Name:     team.Name,
		Initial:  team.Initial,
		LeagueId: team.LeagueId,
	})

}

type GetTeamError struct {
	Message string `json:"message"`
}

type GetTeamSuccess struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initial  string `json:"initial"`
	LeagueId int    `json:"leagueId"`
}
