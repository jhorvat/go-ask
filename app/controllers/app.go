package controllers

import (
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
	bson "labix.org/v2/mgo/bson"
)

type App struct {
	MgoController
}

//Index handler, displays the main question feed
func (c App) Index() revel.Result {
	//Get a mongo session on the questions collection
	col := c.session.DB("go-ask").C("questions")

	//Get a feed of all questions in the collection mapped to the Question model
	var questions []models.Question
	err := col.Find(nil).All(&questions)
	if err != nil {
		panic(err)
	}

	//Pass the questions in the renderargs to be inserted in the template
	return c.Render(questions)
}

func (c App) Question(id string) revel.Result {
	col := c.session.DB("go-ask").C("questions")

	var question models.Question
	err := col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&question)
	if err != nil {
		panic(err)
	}

	return c.Render(question)
}

func (c App) Ask() revel.Result {
	return c.Render()
}
