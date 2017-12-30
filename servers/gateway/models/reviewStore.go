package models

import (
	"gopkg.in/mgo.v2"
	"errors"
)

type ReviewStore struct {
	//the mongo session
	session *mgo.Session
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	col *mgo.Collection
}

func NewMongoStore(sess *mgo.Session, dbName string, colName string) (*ReviewStore, error) {
	d := &ReviewStore{
		session: sess,
		colname: colName,
		dbname: dbName,
		col: sess.DB(dbName).C(colName),
	}

	// TODO: Add indexing for specific values, such as users and whatnot

	return d, nil
}

func (rs *ReviewStore) PostReview(review *Review) (*Review, error) {
	if review == nil {
		return nil, errors.New("review is nil")
	}

	if err := rs.col.Insert(review); err != nil {
		return nil, err
	}

	return review, nil

}