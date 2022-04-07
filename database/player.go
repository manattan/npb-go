package database

import (
	"fmt"

	"github.com/manattan/npb_go/entity"
	"github.com/manattan/npb_go/repository"
	"gorm.io/gorm"
)

var _ repository.Player = &PlayerRepository{}

type PlayerRepository struct {
	DB *gorm.DB
}

type Player struct {
	ID            string
	Name          string
	Year          int
	TeamID        int `gorm:"column:teamId"`
	PositionID    int `gorm:"column:positionId"`
	UniformNumber int `gorm:"column:uniformNumber"`
}

func NewPlayerRepository(DB *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: DB}
}

func (r *PlayerRepository) FindAll() ([]*entity.Player, error) {
	var players []Player

	r.DB.Table("players").Find(&players)

	if len(players) == 0 {
		return nil, fmt.Errorf("player is not registered")
	}

	var res []*entity.Player
	for _, player := range players {
		res = append(res, &entity.Player{
			ID:            player.ID,
			Year:          player.TeamID,
			Name:          player.Name,
			TeamID:        player.TeamID,
			PositionID:    player.PositionID,
			UniformNumber: player.UniformNumber,
		})
	}

	return res, nil
}
