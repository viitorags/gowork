package handler

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/viitorags/gowork/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
    ctx.Header("Content-type", "application/json")
    ctx.JSON(code, gin.H{
        "message":   msg,
        "errorCode": code,
    })
}

func sendSucess(ctx *gin.Context, op string, data interface{}) {
    ctx.Header("Content-type", "application/json")
    ctx.JSON(http.StatusOK, gin.H{
        "message": fmt.Sprintf("operation from handler: %s sucessfully", op),
        "data":    data,
    })
}

// ErrorResponse is the error response
type ErrorResponse struct {
    Message   string `json:"message"`
    ErrorCode string `json:"errorCode"`
}

// CreateWorkResponse is the response for creating a work
type CreateWorkResponse struct {
    Message string                `json:"message"`
    Data    schemas.WorksResponse `json:"data"`
}

// DeleteWorkResponse is the response for deleting a work
type DeleteWorkResponse struct {
    Message string                `json:"message"`
    Data    schemas.WorksResponse `json:"data"`
}

// ShowWorkResponse is the response for showing a work
type ShowWorkResponse struct {
    Message string                `json:"message"`
    Data    schemas.WorksResponse `json:"data"`
}

// GetWorksResponse is the response for listing works
type GetWorksResponse struct {
    Message string                  `json:"message"`
    Data    []schemas.WorksResponse `json:"data"`
}

// UpdateWorkResponse is the response for updating a work
type UpdateWorkResponse struct {
    Message string                `json:"message"`
    Data    schemas.WorksResponse `json:"data"`
}
