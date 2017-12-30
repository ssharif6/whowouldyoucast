package models

import "gopkg.in/mgo.v2/bson"

type Actor struct {
	id bson.ObjectId `json:"id" bson:"_id"`
	firstName string
	lastName string
	// TODO: Add more fields based on the dataset
}


