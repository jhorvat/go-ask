package controllers

import "github.com/robfig/revel"

//Setup request interception
func init() {
	revel.InterceptMethod((*MgoController).Begin, revel.BEFORE)
	revel.InterceptMethod((*MgoController).Close, revel.AFTER)
}
