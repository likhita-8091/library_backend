package router

type Group struct {
	BaseRoute
	UserRoute
	BookKindRoute
}

var BaseGroup = Group{}
