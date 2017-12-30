package models

import "gopkg.in/mgo.v2/bson"

type Suggestion struct {
	id bson.ObjectId `json:"id" bson:"_id"`
}
