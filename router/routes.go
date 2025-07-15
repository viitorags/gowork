package router

import (
    "github.com/gin-gonic/gin"
    "github.com/viitorags/gowork/handler"
)

func InitializeRoutes(router *gin.Engine) {
    v1 := router.Group("/api/v1")
    {
        v1.GET("/works", handler.GetWorksV1)
    }
}
