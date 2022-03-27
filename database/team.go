package database

import (
	"fmt"

	"github.com/manattan/npb_api/entity"
	"github.com/manattan/npb_api/repository"
	"gorm.io/gorm"
)

var _ repository.Team = &TeamRepository{}

type TeamRepository struct {
	DB *gorm.DB
}

type Team struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Initial  string `json:"initial"`
	LeagueId int    `json:"leagueId"`
}

func NewTeamRepository(DB *gorm.DB) *TeamRepository {
	return &TeamRepository{DB: DB}
}

func (r *TeamRepository) FindByID(id string) (*entity.Team, error) {
	var team Team
	r.DB.Table("teams").First(&team, id)
	fmt.Println(team)

	if team.ID == "" || team.Name == "" {
		return nil, fmt.Errorf("team id %s is not found", id)
	}

	return &entity.Team{
		ID:       team.ID,
		Name:     team.Name,
		Initial:  team.Initial,
		LeagueId: team.LeagueId,
	}, nil
}
