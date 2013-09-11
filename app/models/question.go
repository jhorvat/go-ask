package models

import (
	bson "labix.org/v2/mgo/bson"
)

type Question struct {
	Id    bson.ObjectId "_id,omitempty"
	Title string        "title"
	Body  string        "body"
}
