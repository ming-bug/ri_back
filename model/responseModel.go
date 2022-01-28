package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Result struct {
	Ctx *gin.Context
}

// 响应体
type ResultCont struct {
	Success      bool        `json:"success" example:"true"`
	Data         interface{} `json:"data"`
	ErrorCode    string      `json:"errorCode" example:"200"`
	ErrorMessage string      `json:"errorMessage" example:""`
	ShowType     int         `json:"showType" example:"0"` // error display type： 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page
}

type LoginResponse struct {
	Token string `json:"token"`
}

func NewResult(c *gin.Context) *Result {
	return &Result{Ctx: c}
}

// 成功
func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	rc := ResultCont{
		Success: true,
		Data:    data,
	}
	r.Ctx.JSON(http.StatusOK, rc)
}

// 失败
func (r *Result) Faild(code int, msg string, showType int) {
	rc := ResultCont{
		Success:      false,
		Data:         nil,
		ErrorCode:    strconv.Itoa(code),
		ErrorMessage: msg,
		ShowType:     showType,
	}
	r.Ctx.JSON(code, rc)
}
