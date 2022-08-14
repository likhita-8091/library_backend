package auth

import (
	"errors"
	"github.com/CodingJzy/library_backend/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log"
)

func PreReq(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 取jwt token信息
		userAny := c.Get("user")
		token, ok := userAny.(*jwt.Token)
		if !ok {
			return errors.New("token断言失败")
		}

		claims, ok := token.Claims.(*MyJwtClaims)
		if !ok {
			return errors.New("token断言失败")
		}

		// 设置两个变量：表示当前登陆的用户和角色
		c.Set("login_user_id", claims.ID)
		c.Set("login_user", claims.Name)
		c.Set("login_role", claims.Role)
		if model.Role(claims.Role) == model.Admin {
			c.Set("admin", true)
		}
		log.Println("set login info success")
		return next(c)
	}
}
