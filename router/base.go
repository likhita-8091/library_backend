package router

import (
	"github.com/CodingJzy/library_backend/api"
	"github.com/labstack/echo/v4"
)

type BaseRoute struct{}

func (b *BaseRoute) InitBaseRoute(g *echo.Group) {
	g.POST("/login", api.AllGroup.Login)
}
