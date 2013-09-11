package controllers

import (
	r "github.com/robfig/revel"
	mgo "labix.org/v2/mgo"
)

//Wrapper Controller for a Revel Controller that pacakges a MongoDB session with it
type MgoController struct {
	*r.Controller
	session *mgo.Session
}

//Intercept GETs and inject the MongoDB session
func (c *MgoController) Begin() r.Result {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	c.session = session
	return nil
}

//Make sure to close the session when we move on
func (c *MgoController) Close() r.Result {
	c.session.Close()
	return nil
}
