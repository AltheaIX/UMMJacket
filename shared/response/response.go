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

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, ApiResponse{StatusCode: code, Error: msg})
}
