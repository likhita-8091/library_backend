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

func getOperateMsg(c echo.Context) (msg string) {
	method := c.Request().Method
	switch method {
	case http.MethodGet:
		msg = "查询成功"
	case http.MethodPost:
		msg = "添加成功"
	case http.MethodPut:
		msg = "更新成功"
	case http.MethodDelete:
		msg = "删除成功"
	default:
		msg = "操作成功"
	}
	return msg
}

func Ok(c echo.Context) error {
	return Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c echo.Context) error {
	return Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c echo.Context) error {
	return Result(SUCCESS, data, getOperateMsg(c), c)
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
