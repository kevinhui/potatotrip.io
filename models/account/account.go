package account

import (
	"github.com/globalsign/mgo/bson"
	mongorepo "github.com/kevinhui/potatotrip.io/core/repo"
)

type Account struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	Password string        `bson:"password"`
}

var repo, err = mongorepo.NewRepo("localhost")
var db = "potrips"
var table = "account"

func Get(selector bson.M) *Account {
	m := &Account{}
	sess := repo.GetSession()
	defer sess.Close()
	coll := sess.DB(db).C(table)
	if err := coll.Find(selector).One(&m); err != nil {
		return nil
	}
	return m
}
