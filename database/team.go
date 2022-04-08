package database

import (
	"fmt"

	"github.com/manattan/npb_go/entity"
	"github.com/manattan/npb_go/repository"
	"gorm.io/gorm"
)

var _ repository.Team = &TeamRepository{}

type TeamRepository struct {
	DB *gorm.DB
}

type Team struct {
	ID      string
	Name    string
	Initial string
	// 2022-03-27 gorm basic policy is snake_case. However, local DB columns are camelCase
	// TODO: convert camelCase to snake_case
	LeagueID int `gorm:"column:leagueId"`
}

func NewTeamRepository(DB *gorm.DB) *TeamRepository {
	return &TeamRepository{DB: DB}
}

func (r *TeamRepository) FindByID(id string) (*entity.Team, error) {
	var team Team
	r.DB.Table("teams").First(&team, id)

	if team.ID == "" || team.Name == "" {
		return nil, fmt.Errorf("team id %s is not found", id)
	}

	return &entity.Team{
		ID:       team.ID,
		Name:     team.Name,
		Initial:  team.Initial,
		LeagueId: team.LeagueID,
	}, nil
}

func (r *TeamRepository) FindAll() ([]*entity.Team, error) {
	var teams []Team

	r.DB.Table("teams").Find(&teams)

	if len(teams) == 0 {
		return nil, fmt.Errorf("team is not registered")
	}

	var res []*entity.Team
	for _, team := range teams {
		res = append(res, &entity.Team{
			ID:       team.ID,
			Name:     team.Name,
			Initial:  team.Initial,
			LeagueId: team.LeagueID,
		})
	}

	return res, nil
}
