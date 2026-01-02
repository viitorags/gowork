package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viitorags/gowork/schemas"
)

// @Summary        Get works
// @Description    List all job
// @Tags            Works
// @Accept            json
// @Produce        json
// @Success        200    {object}    GetWorksResponse
// @Failure        500    {object}    ErrorResponse
// @Router            /works [get]
func GetWorksHandler(ctx *gin.Context) {
	works := []schemas.Works{}

	err := db.Find(&works).Error
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error getting works")
		return
	}

	sendSucess(ctx, "list works", works)
}
