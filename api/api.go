package api

type group struct {
	BaseApi
	UserApi
}

var AllGroup = new(group)
