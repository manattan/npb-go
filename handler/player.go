package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manattan/npb_go/usecase"
)

type PlayerHandler struct {
	PlayerUC *usecase.PlayerUseCase
}

func NewPlayerHandler(PlayerUC *usecase.PlayerUseCase) *PlayerHandler {
	return &PlayerHandler{PlayerUC: PlayerUC}
}

func (h *PlayerHandler) GetPlayers(c echo.Context) error {
	players, err := h.PlayerUC.GetPlayers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &GetTeamError{
			Message: err.Error(),
		})
	}

	var res []*GetPlayersSuccess
	for _, player := range players {
		res = append(res, &GetPlayersSuccess{
			ID:            player.ID,
			Name:          player.Name,
			Year:          player.Year,
			TeamID:        player.TeamID,
			PositionID:    player.PositionID,
			UniformNumber: player.UniformNumber,
		})
	}

	return c.JSON(http.StatusOK, res)
}

type GetPlayersSuccess struct {
	ID            string `json:"id"`
	Year          int    `json:"year"`
	Name          string `json:"name"`
	TeamID        int    `json:"teamId"`
	PositionID    int    `json:"positionId"`
	UniformNumber int    `json:"uniformNumber"`
}
