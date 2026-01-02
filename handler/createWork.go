package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viitorags/gowork/schemas"
)

// @Summary        Create work
// @Description    Create a new job
// @Tags            Works
// @Accept            json
// @Produce        json
// @Param            request    body        CreateWorkRequest    true    "Request body"
// @Success        200        {object}    CreateWorkResponse
// @Failure        400        {object}    ErrorResponse
// @Failure        500        {object}    ErrorResponse
// @Router            /works [post]
func CreateWorkHandler(ctx *gin.Context) {
	request := CreateWorkRequest{}

	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		logger.Error("validation error: ", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	work := schemas.Works{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	err = db.Create(&work).Error
	if err != nil {
		logger.Error("error creating opening: ", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(ctx, "create work", work)

}
