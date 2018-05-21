package models

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

type Trip struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `bson:"name"`
	FullName    string        `bson:"full_name"`
	Position    interface{}   `bson:"position"`
	Days        int           `bson:"days"`
	StartWeek   int           `bson:"start_week"`
	Plans       interface{}   `bson:"plans"`
	PlaceList   interface{}   `bson:"place_list"`
	PlaceID     string        `bson:"place_id"`
	CountryCode string        `bson:"country_code"`
	Bounds      interface{}   `bson:"bounds"`
	StartDate   time.Time     `bson:"start_date"`
	EndDate     time.Time     `bson:"end_date"`
	Owner       string        `bson:"owner"`
	Members     []string      `bson:"members"`
}
