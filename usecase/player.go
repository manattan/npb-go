package usecase

import (
	"github.com/manattan/npb_go/entity"
	"github.com/manattan/npb_go/repository"
)

type PlayerUseCase struct {
	PlayerRepo repository.Player
}

func NewPlayerUseCase(PlayerRepo repository.Player) *PlayerUseCase {
	return &PlayerUseCase{PlayerRepo: PlayerRepo}
}

func (u *PlayerUseCase) GetPlayers(params repository.GetPlayersSearchParams) ([]*entity.Player, error) {
	players, err := u.PlayerRepo.FindWithParams(params)

	if err != nil {
		return nil, err
	}

	return players, err
}

type GetPlayersSearchParams struct {
	TeamID        int
	Year          int
	UniformNumber int
}
