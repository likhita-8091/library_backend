package main

import (
	"github.com/CodingJzy/library_backend/core"
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/initialize"
	"github.com/CodingJzy/library_backend/middlewars/auth"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

func userInfo1(c echo.Context) error {
	return c.String(200, " 查看用户信息")
}

func userInfo2(c echo.Context) error {
	return c.String(200, "Welcome "+"1"+"!")
}

type User struct {
	UserName string
	PassWord string
}

// login：签发一个jwt token
func Login(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	if u.UserName == "jiang_wei" && u.PassWord == "echo" {

		// 实例化一个自定义claims
		claims := &auth.MyJwtClaims{
			Name:  "jiang_wei",
			Admin: "true",
			StandardClaims: jwt.StandardClaims{
				IssuedAt: time.Now().Unix(),
				// 过期时间
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// 签发token
		t, err := claims.CreateToken()
		if err != nil {
			return err
		}
		return c.JSON(200, echo.Map{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

func main() {
	// 初始化配置
	core.LoadConfig()

	// 连接数据库
	global.DB = initialize.Mysql()

	// 实例化echo对象
	e := echo.New()

	// 日志中间件、异常捕获中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 登陆
	e.POST("/login", Login)

	// 实例化jwt配置
	jwtConf := middleware.JWTConfig{
		SigningKey: []byte(auth.Secret),
		Claims:     &auth.MyJwtClaims{},
	}

	// 接口
	g := e.Group("/api/v1", middleware.JWTWithConfig(jwtConf))

	g.GET("/user_info1", userInfo1)
	g.GET("/user_info2", userInfo2)

	// 启动server
	e.Logger.Fatal(e.Start(":1323"))
}
