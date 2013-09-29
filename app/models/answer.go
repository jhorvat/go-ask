package models

import (
	bson "labix.org/v2/mgo/bson"
)

//An answer model
type Answer struct {
	Id       bson.ObjectId "_id,omitempty"
	ParentId bson.ObjectId
	Body     string "body"
	Accepted bool   "accepted"
	Points   int    "points"
}
