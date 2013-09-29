package controllers

import (
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
	bson "labix.org/v2/mgo/bson"
)

//Setup request interception and template function
func init() {
	revel.InterceptMethod((*MgoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MgoController).Close, revel.AFTER)

	//Extract the hex ID of a mongo id in our template
	revel.TemplateFuncs["bsonID"] = func(a bson.ObjectId) string { return a.Hex() }
	//Return the len of the answers slice passed
	revel.TemplateFuncs["answerCount"] = func(a []models.Answer) int { return len(a) }
}
