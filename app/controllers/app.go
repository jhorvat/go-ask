package controllers

import (
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
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
	err := col.Find(nil).Sort("-_id").Limit(50).All(&questions)
	if err != nil {
		panic(err)
	}

	//Pass the questions in the renderargs to be inserted in the template
	return c.Render(questions)
}
