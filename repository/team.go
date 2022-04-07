package repository

import "github.com/manattan/npb_go/entity"

type Team interface {
	FindByID(id string) (*entity.Team, error)
	FindAll() ([]*entity.Team, error)
}
