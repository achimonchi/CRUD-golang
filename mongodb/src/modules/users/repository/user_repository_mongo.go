package repository

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/achimonchi/CRUD-golang/mongodb/src/modules/users/model"
)

type userRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

func NewUserRepositoryMongo(db *mgo.Database, collection string) *userRepositoryMongo {
	return &userRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *userRepositoryMongo) Save(user *model.User) error {
	err := r.db.C(r.collection).Insert(user)
	return err
}

func (r *userRepositoryMongo) Update(id string, user *model.User) error {
	user.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"_id": id}, user)
	return err
}

func (r *userRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"_id": id})
	return err
}

func (r *userRepositoryMongo) FindAll() (*model.Users, error) {
	var profiles model.Profiles
	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *userRepositoryMongo) FindByID(id string) (*model.User, error) {
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"_id": id}).One(&profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
