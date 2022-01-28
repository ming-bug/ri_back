package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	_ "ri/docs"
	"ri/router"
)

// @title ri 后台查询
// @version 1.0
// @description Query about summary report of campaign.
// @termsOfService http://swagger.io/terms/
// @contact.name Shimix
// @contact.url
// @contact.email Shiming.Xue@hgc.com.hk
// @license.name
// @license.url
// @host localhost:8002
// @BasePath
func main() {
	r := router.Router()
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8002"); err != nil {
		log.Fatal(err)
	}
}
