package repository

import "github.com/manattan/npb_api/entity"

type Team interface {
	FindByID(id string) (*entity.Team, error)
}
