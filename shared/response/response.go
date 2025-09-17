package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Metadata   interface{} `json:"metadata,omitempty"`
}

func JSON(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, ApiResponse{StatusCode: code, Data: data})
}

func Metadata(ctx *gin.Context, data interface{}, metadata interface{}) {
	ctx.JSON(http.StatusOK, ApiResponse{StatusCode: http.StatusOK, Data: data, Metadata: metadata})
}

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, ApiResponse{StatusCode: code, Error: msg})
}
