package models

import "github.com/globalsign/mgo/bson"

type Account struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}
