package api

import (
	"context"
	"fmt"
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/model"
	"github.com/CodingJzy/library_backend/model/req"
	"github.com/CodingJzy/library_backend/model/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"strconv"
)

type UserApi struct {
}

// 添加用户
func (u *UserApi) AddUser(c echo.Context) error {
	// 解析参数
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if user.Phone != "" && len(user.Phone) != 11 {
		return response.FailWithMessage("手机号格式不正确", c)
	}

	// 如果是添加管理员，禁止添加
	if user.Role == model.Admin {
		return response.FailWithMessage("禁止添加管理员", c)
	}

	// 如果是添加图书管理员，登陆的角色是超级管理员
	currentRole := c.Get("login_role").(int)
	if user.Role == model.BookManager && model.Role(currentRole) != model.Admin {
		return response.FailWithMessage("权限不够", c)
	}

	// 禁止添加新用户时设置密码
	if user.Password != "" {
		return response.FailWithMessage("禁止添加新用户时设置密码", c)
	}

	// 字段校验
	valid := validator.New()
	err = valid.Struct(user)
	if err != nil {
		return response.FailWithMessage(fmt.Sprintf("字段校验失败：%v", err.Error()), c)
	}

	// 给图书管理员设置默认密码：admin
	if user.Role == model.BookManager {
		user.Password = model.EncryptPassword("admin")
	}

	// 存库
	if err = global.DB.Create(user).Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(user, c)
}

// 删除用户
func (u *UserApi) DeleteUser(c echo.Context) error {
	// 解析参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}

	// 查找用户
	var user model.User
	user.ID = uint(id)

	if err = global.DB.First(&user).Debug().Error; err != nil {
		global.DB.Logger.Warn(context.Background(), err.Error())
		return response.FailWithMessage(err.Error(), c)
	}

	// 如果是超级管理员，不能删除
	if user.Role == model.Admin {
		return response.FailWithMessage("超级管理员不允许删除", c)
	}

	// 删除
	if err = global.DB.Where("role != ?", model.Admin).Delete(&user).Debug().Error; err != nil {
		global.DB.Logger.Warn(context.Background(), err.Error())
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(user, c)
}

// 修改用户
func (u *UserApi) UpdateUser(c echo.Context) error {
	// 解析参数
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if user.Name == "" {
		return response.FailWithMessage("用户名不能为空", c)
	}

	if user.Phone != "" && len(user.Phone) != 11 {
		return response.FailWithMessage("手机号格式不正确", c)
	}

	// 禁止编辑用户时设置密码
	if user.Password != "" {
		return response.FailWithMessage("禁止编辑用户时设置密码", c)
	}

	if user.Role != 0 {
		return response.FailWithMessage("禁止编辑用户时设置角色", c)
	}

	// 取要修改的用户信息ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}

	// 根据ID查找用户
	var beforeUser model.User
	beforeUser.ID = uint(id)
	err = global.DB.First(&beforeUser).Debug().Error
	if err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	// 存库
	if err := global.DB.Model(beforeUser).Updates(user).Find(user).Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}
	return response.OkWithData(user, c)
}

// 获取所有用户
func (u *UserApi) ListUser(c echo.Context) error {
	var userList []model.User

	currentRole := c.Get("login_role").(int)

	// 超级管理员可以获取所有用户
	if model.Role(currentRole) == model.Admin {
		if err := global.DB.Find(&userList).Debug().Error; err != nil {
			return response.FailWithMessage(err.Error(), c)
		}
	} else {
		// 图书管理员只能获取所有读者
		if err := global.DB.Where("role = ?", model.Reader).Find(&userList).Debug().Error; err != nil {
			return response.FailWithMessage(err.Error(), c)
		}
	}

	return response.OkWithData(userList, c)
}

// 获取单个用户
func (u *UserApi) GetUser(c echo.Context) error {
	// 解析参数
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}
	var user model.User
	user.ID = uint(id)

	// 查询
	if err = global.DB.First(&user).Debug().Error; err != nil {
		global.DB.Logger.Warn(context.Background(), err.Error())
		return response.FailWithMessage(err.Error(), c)
	}

	return response.OkWithData(user, c)
}

// 修改密码
func (u UserApi) ChangePassword(c echo.Context) error {
	// 如果是超级管理员，可以重置图书管理员密码
	manageID := c.QueryParam("book_manage_id")

	if manageID != "" {
		isAdmin := c.Get("admin")
		_, ok := isAdmin.(bool)
		if ok {
			var user model.User
			if err := global.DB.Where("id = ?", manageID).First(&user).Error; err != nil {
				return response.FailWithMessage(err.Error(), c)
			}
			user.Password = model.EncryptPassword("admin")
			if err := global.DB.Save(user).Error; err != nil {
				return response.FailWithMessage(err.Error(), c)
			}
			return response.OkWithMessage(fmt.Sprintf("重置%v密码成功", user.Name), c)
		}
		return response.FailWithMessage("权限不够", c)
	}

	// 参数解析
	param := req.ChangePasswordReq{}
	err := c.Bind(&param)
	if err != nil {
		return response.FailWithMessage("parse param error", c)
	}

	// 取出当前登陆用户的ID
	userID := c.Get("login_user_id").(uint)

	// 数据库查询该用户的旧密码与提交的信息比较，如果一致允许修改
	var user model.User
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return response.FailWithMessage(err.Error(), c)
	}

	if !model.DecryptPassword(param.Password, user.Password) {
		return response.FailWithDetailed(param.Password, "提交的旧密码不正确", c)
	}

	// 新密码赋值
	user.Password = model.EncryptPassword(param.NewPassword)
	user.FirstLogin = false
	if err := global.DB.Save(&user).Error; err != nil {
		return response.FailWithMessage("密码修改失败", c)
	}

	return response.OkWithMessage("密码修改成功", c)
}
