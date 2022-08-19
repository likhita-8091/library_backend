package auth

import (
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/model"
	"github.com/CodingJzy/library_backend/model/response"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

func CheckID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return next(c)
		}

		userID, err := strconv.Atoi(id)
		if err != nil {
			return response.FailWithMessage("parse param error", c)
		}

		var user model.User
		user.ID = uint(userID)

		// 查询该ID的用户信息
		if err = global.DB.First(&user).Debug().Error; err != nil {
			return response.FailWithMessage(err.Error(), c)
		}

		// 获取当前登陆用的的角色、ID
		currentRole := c.Get("login_role").(int)
		currentUserID := c.Get("login_user_id").(uint)

		// 当前登陆的用户为图书管理员，操作的用户角色为超级管理员，禁止修改
		if model.Role(currentRole) == model.BookManager && user.Role == model.Admin {
			return response.FailWithMessage("禁止下级越权", c)
		}

		// 当前登陆的用户为图书管理员，但是操作的用户角色同样为其余图书管理员，禁止修改
		if model.Role(currentRole) == model.BookManager && user.Role == model.BookManager {
			if currentUserID != user.ID {
				return response.FailWithMessage("禁止同级越权", c)
			}
		}

		// 如果是删除用户，不能删除超级管理员、不能删除自身
		if strings.EqualFold("delete", c.Request().Method) {
			if user.Role == model.Admin {
				return response.FailWithMessage("超级管理员无法被删除", c)
			}
			if currentUserID == user.ID {
				return response.FailWithMessage("禁止删除自身", c)
			}
		}

		return next(c)
	}
}
