package api

import (
	"context"
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/model"
	"github.com/CodingJzy/library_backend/model/response"
	"github.com/labstack/echo/v4"

	"strconv"
)

type BookKindApi struct {
}

// 添加图书分类
func (u *BookKindApi) AddBookKind(c echo.Context) error {
	// 解析参数
	bookKind := new(model.BookKind)
	err := c.Bind(bookKind)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if bookKind.Name == "" {
		return response.FailWithMessage("图书分类名不能为空", c)
	}

	// 存库
	if err = global.DB.Create(bookKind).Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(bookKind, c)
}

// 删除图书分类
func (u *BookKindApi) DeleteBookKind(c echo.Context) error {
	// 解析参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}

	// 删除
	var bookKind model.BookKind
	bookKind.ID = uint(id)
	if err = global.DB.Delete(&bookKind).Debug().Error; err != nil {
		global.DB.Logger.Warn(context.Background(), err.Error())
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithMessage("删除成功", c)
}

// 修改图书分类
func (u *BookKindApi) UpdateBookKind(c echo.Context) error {
	// 解析参数
	bookKind := new(model.BookKind)
	err := c.Bind(bookKind)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if bookKind.Name == "" {
		return response.FailWithMessage("图书分类名称不能为空", c)
	}

	// 取要修改的图书分类信息ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}

	// 根据ID查找图书分类
	var beforeBookKind model.BookKind
	beforeBookKind.ID = uint(id)
	err = global.DB.First(&beforeBookKind).Debug().Error
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	// 存库
	beforeBookKind.Name = bookKind.Name

	if err := global.DB.Save(&beforeBookKind).Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	return response.OkWithData(beforeBookKind, c)
}

// 获取所有图书分类
func (u *BookKindApi) ListBookKind(c echo.Context) error {
	var bookKindList []model.BookKind

	if err := global.DB.Find(&bookKindList).Debug().Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(bookKindList, c)
}

// 获取单个图书分类
func (u *BookKindApi) GetbookKind(c echo.Context) error {
	// 解析参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}
	var bookKind model.BookKind
	bookKind.ID = uint(id)

	// 查询
	if err = global.DB.First(&bookKind).Debug().Error; err != nil {
		global.DB.Logger.Warn(context.Background(), err.Error())
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(bookKind, c)
}
