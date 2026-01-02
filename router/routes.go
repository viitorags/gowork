package router

import (
	"net/http"

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

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	const worksPath = "/works"
	const workByIDPath = "/works/:id"

	v1 := router.Group(basePath)
	{
		v1.GET(worksPath, handler.GetWorksHandler)
		v1.GET(workByIDPath, handler.ShowWorkHandler)
		v1.POST(worksPath, handler.CreateWorkHandler)
		v1.DELETE(workByIDPath, handler.DeleteWorkHandler)
		v1.PUT(workByIDPath, handler.UpdateWorkHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
