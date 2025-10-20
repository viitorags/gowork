package router

import (
    "github.com/gin-gonic/gin"
)

func Initialize() {
    router := gin.Default()

    InitializeRoutes(router)

    router.Run(":8080")
}
