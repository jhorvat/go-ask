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

	return c.Render(question)
}

func (c Questions) Ask() revel.Result {
	return c.Render()
}

func (c Questions) Submit(title, body, tags string) revel.Result {
	col := c.session.DB("go-ask").C("questions")

	err := col.
	return c.Redirect(Questions.Ask)
}
