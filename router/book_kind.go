package router

import (
	"github.com/CodingJzy/library_backend/api"
	"github.com/labstack/echo/v4"
)

type BookKindRoute struct{}

func (b *BookKindRoute) InitBookKindRoute(g *echo.Group) {
	g.POST("/book_kinds", api.AllGroup.AddBookKind)
	g.DELETE("/book_kinds/:id", api.AllGroup.DeleteBookKind)
	g.PUT("/book_kinds/:id", api.AllGroup.UpdateBookKind)
	g.GET("/book_kinds", api.AllGroup.ListBookKind)
	g.GET("/book_kinds/:id", api.AllGroup.GetbookKind)
}
