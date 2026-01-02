package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viitorags/gowork/schemas"
)

// @Summary        Update work
// @Description    Update a job
// @Tags            Works
// @Accept            json
// @Produce        json
// @Param            id        path        string                true    "Work identification"
// @Param            request    body        UpdateWorkRequest    true    "Work data to Update"
// @Success        200        {object}    UpdateWorkResponse
// @Failure        400        {object}    ErrorResponse
// @Failure        404        {object}    ErrorResponse
// @Failure        500        {object}    ErrorResponse
// @Router            /works/{id} [put]
func UpdateWorkHandler(ctx *gin.Context) {
	request := UpdateWorkRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Error("validation error: ", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	work := schemas.Works{}

	if err := db.First(&work, "id = ?", id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "work not found")
		return
	}

	if request.Role != "" {
		work.Role = request.Role
	}

	if request.Company != "" {
		work.Company = request.Company
	}

	if request.Location != "" {
		work.Location = request.Location
	}

	if request.Remote != nil {
		work.Remote = *request.Remote
	}

	if request.Link != "" {
		work.Link = request.Link
	}

	if request.Salary > 0 {
		work.Salary = request.Salary
	}

	if err := db.Save(&work).Error; err != nil {
		logger.Error("error updating work: ", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating work")
		return
	}

	sendSucess(ctx, "work update", work)
}
