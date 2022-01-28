package router

import (
	"github.com/gin-gonic/gin"
	"ri/handler"
	"ri/handler/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	// 处理异常
	router.NoRoute(middleware.NotFound)
	router.NoMethod(middleware.NotFound)
	router.Use(middleware.Recover)
	// router
	user := router.Group("/api/user")
	user.GET("/current", handler.CurrentUser)
	user.GET("/groups", handler.AllGroups)
	user.Use(middleware.Auth())
	{
		user.POST("/login", handler.Login)
		user.POST("/outLogin", handler.OutLogin)
	}

	campaign := router.Group("/api/campaign")
	campaign.GET("/myCampaigns", handler.MyCampaigns)
	router.Group("/api/list")
	return router
}
