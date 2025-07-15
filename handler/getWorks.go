package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetWorksV1(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "msg": "GET WORKS",
    })
}
