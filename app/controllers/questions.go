package controllers

import (
	"fmt"
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
	bson "labix.org/v2/mgo/bson"
	"strings"
)

type Questions struct {
	App
}

func (c Questions) Show(id string) revel.Result {
	col := c.session.DB("go-ask").C("questions")

	var question models.Question
	err := col.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&question)
	if err != nil {
		panic(err)
	}

	col = c.session.DB("go-ask").C("answers")

	answers := []models.Answer{}
	err = col.Find(bson.M{"ParentId": bson.ObjectIdHex(id)}).All(&answers)
	if err != nil {
		panic(err)
	}

	for _, answer := range answers {
		fmt.Printf("%s\n", answer.Body)
	}

	return c.Render(question)
}

func (c Questions) Ask() revel.Result {
	return c.Render()
}

func (c Questions) Submit(title, body, tags string) revel.Result {
	id := bson.NewObjectId()
	col := c.session.DB("go-ask").C("questions")

	err := col.Insert(&models.Question{id, title, body, strings.Split(tags, ",")})
	if err != nil {
		panic(err)
	}
	return c.Redirect("/question/%s", id.Hex())
}

func (c Questions) Answer(id, body string) revel.Result {
	answerId := bson.NewObjectId()
	col := c.session.DB("go-ask").C("answers")

	err := col.Insert(&models.Answer{answerId, bson.ObjectIdHex(id), body})
	if err != nil {
		panic(err)
	}

	return c.Redirect("/question/%s", id)
}
