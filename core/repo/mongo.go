package repo

import (
	"fmt"

	"github.com/globalsign/mgo"
)

type Repo struct {
	session *mgo.Session
}

func NewRepo(durl string) (*Repo, error) {
	session, err := mgo.Dial("mongodb://" + durl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return &Repo{
		session,
	}, nil
}

func (r *Repo) GetSession() *mgo.Session {
	return r.session.Clone()
}

func (r *Repo) Close() {
	r.session.Close()
	return
}
