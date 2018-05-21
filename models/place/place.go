package place

import "github.com/globalsign/mgo/bson"

type Place struct {
	ID            bson.ObjectId `bson:"_id"`
	PlaceID       string        `bson:"place_id"`
	PageToken     string        `bson:"page_token"`
	NextPageToken string        `bson:"next_page_token"`
	Results       interface{}   `bson:"results"`
}
