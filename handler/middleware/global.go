package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ri/enum"
	"ri/model"
	"ri/utils"
	"runtime/debug"
)

// 认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := model.NewResult(c)
		path := c.FullPath()
		if path == "/api/user/login" {
			c.Next()
			return
		}
		auth := c.GetHeader("x-access-token")
		if auth == "" {
			result.Faild(http.StatusUnauthorized, "No authorization", enum.Error)
			c.Abort()
			return
		}
		token, claims, err := utils.ParseToken(auth)
		if err != nil || !token.Valid {
			result.Faild(http.StatusInternalServerError, "error of parsing token!", enum.Page)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

// 404
func NotFound(c *gin.Context) {
	model.NewResult(c).Faild(http.StatusNotFound, "resource not found", enum.Page)
	return
}

// 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			model.NewResult(c).Faild(http.StatusInternalServerError, "internal server error", enum.Page)
			c.Abort()
		}
	}()
	// 继续后续接口调用
	c.Next()
}
