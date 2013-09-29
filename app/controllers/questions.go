package controllers

import (
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
	bson "labix.org/v2/mgo/bson"
	"strings"
)

type Questions struct {
	App
}

//Displays an individual question
func (c Questions) Show(id string) revel.Result {
	col := c.session.DB("go-ask").C("questions")

	//Get the question from mongo
	var question models.Question
	err := col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&question)
	if err != nil {
		panic(err)
	}

	//Get the answers for the question based on question ID
	col = c.session.DB("go-ask").C("answers")
	answers := []models.Answer{}
	err = col.Find(bson.M{"parentid": bson.ObjectIdHex(id)}).All(&answers)
	if err != nil {
		panic(err)
	}

	return c.Render(question, answers)
}

//Render the question submission form
func (c Questions) Ask() revel.Result {
	return c.Render()
}

//Submit a new question
func (c Questions) Submit(title, body, tags string) revel.Result {
	//Generate a new question ID and insert it into the questions collection
	id := bson.NewObjectId()
	col := c.session.DB("go-ask").C("questions")

	err := col.Insert(&models.Question{id, title, body, strings.Split(tags, ","), 0})
	if err != nil {
		panic(err)
	}

	//Redirect to the new question's page
	return c.Redirect("/question/%s", id.Hex())
}

//Handle an answer POST
func (c Questions) Answer(id, body string) revel.Result {
	//Generate a new answer ID and insert it into the answers collection
	answerId := bson.NewObjectId()
	col := c.session.DB("go-ask").C("answers")

	err := col.Insert(&models.Answer{answerId, bson.ObjectIdHex(id), body, false, 0})
	if err != nil {
		panic(err)
	}

	//Redirect back to the question page this time displaying the new answer
	return c.Redirect("/question/%s", id)
}

func (c Questions) Up(id string) revel.Result {
	col := c.session.DB("go-ask").C("questions")

	//Get the question from mongo
	var question models.Question
	err := col.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$inc": {"points": +1}})
	if err != nil {
		panic(err)
	}

	return c.Redirect("/question/%s", id)
}
