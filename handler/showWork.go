package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/viitorags/gowork/schemas"
)

// @Summary        Show work
// @Description    Show a job
// @Tags            Works
// @Accept            json
// @Produce        json
// @Param            id    path        string    true    "Work identification"
// @Success        200    {object}    ShowWorkResponse
// @Failure        400    {object}    ErrorResponse
// @Failure        404    {object}    ErrorResponse
// @Router            /works/{id} [get]
func ShowWorkHandler(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
        return
    }

    work := schemas.Works{}

    if err := db.First(&work, id).Error; err != nil {
        sendError(ctx, http.StatusNotFound, fmt.Sprintf("work: %s not found", id))
        return
    }

    sendSucess(ctx, "find work", work)
}
