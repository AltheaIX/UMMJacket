package middleware

import (
	AuthServices "github.com/AltheaIX/UMMJacket/internal/domain/auth/service"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	GetUserInfo(c *gin.Context)
}

type AuthMiddlewareImpl struct {
	authServices AuthServices.AuthServices
}

func NewAuthMiddleware(authServices AuthServices.AuthServices) AuthMiddleware {
	return &AuthMiddlewareImpl{authServices: authServices}
}

func (a *AuthMiddlewareImpl) GetUserInfo(c *gin.Context) {
	authorization := c.GetHeader("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	authorizationSplit := strings.Split(authorization, " ")
	if len(authorizationSplit) != 2 {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	authorizationToken := authorizationSplit[1]
	userInfo, err := a.authServices.AuthenticateToken(c, authorizationToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	c.Set("userInfo", userInfo)
	c.Next()
}
