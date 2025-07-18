package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/viitorags/gowork/docs"
	"github.com/viitorags/gowork/handler"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/works", handler.GetWorksHandler)
		v1.GET("/works/:id", handler.ShowWorkHandler)
		v1.POST("/works", handler.CreateWorkHandler)
		v1.DELETE("/works/:id", handler.DeleteWorkHandler)
		v1.PUT("/works/:id", handler.UpdateWorkHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
