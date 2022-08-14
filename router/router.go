package router

type Group struct {
	BaseRoute
	UserRoute
}

var BaseGroup = Group{}
