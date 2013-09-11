package controllers

import (
	"fmt"
	"github.com/jhorvat/go-ask/app/models"
	"github.com/robfig/revel"
)

type App struct {
	MgoController
}

func (c App) Index() revel.Result {
	col := c.session.DB("go-ask").C("questions")
	var questions []models.Question

	err := col.Find(nil).All(&questions)
	if err != nil {
		panic(err)
	}

	for _, q := range questions {
		fmt.Printf("ID: %s\n\t%s\n\t%s\n", q.Id, q.Title, q.Body)
	}
	return c.Render(questions)
}
