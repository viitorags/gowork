package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viitorags/gowork/schemas"
)

// @Summary        Delete work
// @Description    Delete a job
// @Tags            Works
// @Accept            json
// @Produce        json
// @Param            id    path        string    true    "Work identification"
// @Success        200    {object}    DeleteWorkResponse
// @Failure        400    {object}    ErrorResponse
// @Failure        404    {object}    ErrorResponse
// @Router            /works/{id} [delete]
func DeleteWorkHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	work := schemas.Works{}

	if err := db.First(&work, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("work: %s not found", id))
		return
	}

	if err := db.Delete(&work).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSucess(ctx, "delete work", work)
}
