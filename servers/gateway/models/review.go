package models

import "gopkg.in/mgo.v2/bson"

type Review struct {
	id          bson.ObjectId `json:"id" bson:"_id"`
	numStars    int           `json:"numStars"`
	description string        `json:"description"`
}
