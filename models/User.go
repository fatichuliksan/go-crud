package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password,omitempty" bson:"password"`
		Name     string        `json:"name,omitempty" bson:"name"`
	}
)

func (u User) getCollectionName() string {
	return "users"
}
