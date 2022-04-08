package repository

import (
	"github.com/manattan/npb_go/entity"
)

type Player interface {
	FindWithParams(params GetPlayersSearchParams) ([]*entity.Player, error)
}

type GetPlayersSearchParams struct {
	TeamID        int
	Year          int
	UniformNumber int
}
