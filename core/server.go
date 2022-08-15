package core

import (
	"github.com/CodingJzy/library_backend/global"
	"github.com/CodingJzy/library_backend/middlewars/auth"
	"github.com/CodingJzy/library_backend/model/response"
	"github.com/CodingJzy/library_backend/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunServer() {

	// 实例化echo对象
	e := echo.New()

	// 日志中间件、异常捕获中间件、跨域
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	baseGroup := e.Group("")

	{
		// login
		router.BaseGroup.InitBaseRoute(baseGroup)
	}

	{
		// 实例化jwt配置
		jwtConf := middleware.JWTConfig{
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return response.FailWithDetailed(err.Error(), "未认证", c)
			},
			SigningKey: []byte(auth.Secret),
			Claims:     &auth.MyJwtClaims{},
		}
		// 该组下的路由需要jwt 认证
		APIV1Group := baseGroup.Group("/api/v1", middleware.JWTWithConfig(jwtConf), auth.PreReq)
		router.BaseGroup.InitUserRouteRoute(APIV1Group)
	}

	// 启动server
	e.Logger.Fatal(e.Start(global.Config.System.Addr))
}
