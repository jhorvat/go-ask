// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/robfig/revel"


type tMgoController struct {}
var MgoController tMgoController


func (_ tMgoController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MgoController.Begin", args).Url
}

func (_ tMgoController) Close(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("MgoController.Close", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tQuestions struct {}
var Questions tQuestions


func (_ tQuestions) Show(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Questions.Show", args).Url
}

func (_ tQuestions) Ask(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Questions.Ask", args).Url
}

func (_ tQuestions) Submit(
		title string,
		body string,
		tags string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "title", title)
	revel.Unbind(args, "body", body)
	revel.Unbind(args, "tags", tags)
	return revel.MainRouter.Reverse("Questions.Submit", args).Url
}

func (_ tQuestions) Answer(
		id string,
		body string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	revel.Unbind(args, "body", body)
	return revel.MainRouter.Reverse("Questions.Answer", args).Url
}

func (_ tQuestions) Up(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Questions.Up", args).Url
}


