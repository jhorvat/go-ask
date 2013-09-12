package models

import (
	bson "labix.org/v2/mgo/bson"
)

//A question model that contains all of the relevant fields for a question
type Question struct {
	Id    bson.ObjectId "_id,omitempty"
	Title string        "title"
	Body  string        "body"
	Tags  []string      "tags"
}
