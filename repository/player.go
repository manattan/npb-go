package repository

import "github.com/manattan/npb_go/entity"

type Player interface {
	FindAll() ([]*entity.Player, error)
}
