package repo

import (
	"fmt"

	"github.com/globalsign/mgo"
)

type Repo struct {
	session *mgo.Session
}

type Collection struct {
	*mgo.Collection
}

func NewRepo(durl string) (*Repo, error) {
	sess, err := mgo.Dial("mongodb://" + durl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sess.SetMode(mgo.Monotonic, true)
	return &Repo{
		sess,
	}, nil
}

func (r *Repo) GetSession() *mgo.Session {
	return r.session.Clone()
}

func (r *Repo) GetCollection(db, table string) *Collection {
	sess := r.GetSession()
	return &Collection{sess.DB(db).C(table)}
}

func (c *Collection) Close() {
	c.Database.Session.Close()
}
