package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c echo.Context) error {
	return c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

const (
	ERROR   = 0
	SUCCESS = 1
)

func Ok(c echo.Context) error {
	return Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c echo.Context) error {
	return Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c echo.Context) error {
	return Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data interface{}, message string, c echo.Context) error {
	return Result(SUCCESS, data, message, c)
}

func Fail(c echo.Context) error {
	return Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c echo.Context) error {
	return Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c echo.Context) error {
	return Result(ERROR, data, message, c)
}
