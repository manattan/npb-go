package usecase

import (
	"github.com/manattan/npb_api/entity"
	"github.com/manattan/npb_api/repository"
)

type TeamUseCase struct {
	TeamRepo repository.Team
}

func NewTeamUseCase(TeamRepo repository.Team) *TeamUseCase {
	return &TeamUseCase{TeamRepo: TeamRepo}
}

func (u *TeamUseCase) GetTeam(teamId string) ( *entity.Team, error ) {
	team, err := u.TeamRepo.FindByID(teamId)

	if err != nil {
		return nil, err
	}

	return team, nil
}