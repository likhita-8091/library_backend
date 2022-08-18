package router

import (
	"github.com/CodingJzy/library_backend/api"
	"github.com/CodingJzy/library_backend/middlewars/auth"
	"github.com/labstack/echo/v4"
)

type UserRoute struct{}

func (b *UserRoute) InitUserRoute(g *echo.Group) {
	g.POST("/users", api.AllGroup.AddUser)
	g.DELETE("/users/:id", api.AllGroup.DeleteUser, auth.CheckID)
	g.PUT("/users/:id", api.AllGroup.UpdateUser, auth.CheckID)
	g.GET("/users", api.AllGroup.ListUser)
	g.GET("/users/:id", api.AllGroup.GetUser, auth.CheckID)
	g.POST("/users/change_password", api.AllGroup.ChangePassword)
}
