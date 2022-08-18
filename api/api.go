package api

type group struct {
	BaseApi
	UserApi
	BookKindApi
}

var AllGroup = new(group)
