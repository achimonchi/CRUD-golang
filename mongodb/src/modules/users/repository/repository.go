package repository

import (
	"github.com/achimonchi/CRUD-golang/mongodb/src/modules/users/model"
)

type UserRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindAll() (*model.Profiles, error)
	FindByID(string) (*model.Profile, error)
}
