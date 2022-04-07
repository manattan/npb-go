package usecase

import (
	"github.com/manattan/npb_go/entity"
	"github.com/manattan/npb_go/repository"
)

type TeamUseCase struct {
	TeamRepo repository.Team
}

func NewTeamUseCase(TeamRepo repository.Team) *TeamUseCase {
	return &TeamUseCase{TeamRepo: TeamRepo}
}

func (u *TeamUseCase) GetTeam(teamId string) (*entity.Team, error) {
	team, err := u.TeamRepo.FindByID(teamId)

	if err != nil {
		return nil, err
	}

	return team, nil
}

func (u *TeamUseCase) GetTeams() ([]*entity.Team, error) {
	teams, err := u.TeamRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return teams, err
}
