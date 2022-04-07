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

func (u *PlayerUseCase) GetPlayers() ([]*entity.Player, error) {
	players, err := u.PlayerRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return players, err
}
