package api

import (
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/middlewars/auth"
	"github.com/CodingJzy/library_backend/model"
	"github.com/CodingJzy/library_backend/model/req"
	"github.com/CodingJzy/library_backend/model/response"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type BaseApi struct {
}

// login：签发一个jwt token
func (b *BaseApi) Login(c echo.Context) error {
	u := new(req.Login)
	if err := c.Bind(u); err != nil {
		return err
	}

	// 校验
	if u.Name == "" || u.Password == "" {
		return response.FailWithMessage("用户名或密码不能为空", c)
	}

	var user *model.User

	err := global.DB.Where("name = ? ", u.Name).First(&user).Debug().Error
	if err != nil {
		return response.FailWithMessage("用户名或密码不正确", c)
	}

	if !model.DecryptPassword(u.Password, user.Password) {
		return response.FailWithMessage("用户名或密码不正确", c)
	}

	if user.Role == model.Reader {
		return response.FailWithMessage("读者禁止登陆", c)
	}

	// 实例化一个自定义claims
	claims := &auth.MyJwtClaims{
		ID:   user.ID,
		Name: user.Name,
		Role: int(user.Role),
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			// 过期时间
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// 签发token
	t, err := claims.CreateToken()
	if err != nil {
		return response.FailWithMessage("签发token失败", c)
	}

	log.Println("login success ok", user.Name)

	data := echo.Map{
		"is_admin":    false,
		"token":       t,
		"first_login": user.FirstLogin,
	}
	if user.Role == model.Admin {
		data["is_admin"] = true
	}

	return response.OkWithDetailed(data, "登陆成功", c)
}
