package handlers

import (
	"errors"
	AuthModel "github.com/AltheaIX/UMMJacket/internal/domain/auth/model"
	"github.com/AltheaIX/UMMJacket/internal/domain/auth/model/dto"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

// Login godoc
// @Summary      Login Endpoint
// @Description  Api for Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Payload"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/auth/login [post]
func (h *Handlers) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[Login][ShouldBindJSON] ")
		err = &shared.AppError{Code: http.StatusBadRequest, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	validator := shared.GetValidator()
	if err := validator.Struct(request); err != nil {
		log.Error().Err(err).Msg("[Login][Validator] ")
		err = &shared.AppError{Code: http.StatusBadRequest, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	accessToken, refreshToken, err := h.authService.Login(c, request.Nim, request.Password)
	if err != nil {
		code := shared.GetCode(err)

		if code != 200 {
			log.Error().Err(err).Msg("[Login][LoginService] ")
		}

		response.Error(c, code, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

// Refresh godoc
// @Summary      Refresh Endpoint
// @Description  Api for Refresh Token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/auth/refresh [post]
func (h *Handlers) Refresh(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	user := userInfo.(*AuthModel.Claims)

	accessToken, err := h.authService.Refresh(c, user)
	if err != nil {
		log.Error().Err(err).Msg("[Refresh][RefreshService] ")
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, &dto.RefreshResponse{AccessToken: accessToken})
}

// CurrentUser godoc
// @Summary      Current User Endpoint
// @Description  Api for Getting Current User
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Success      200  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/auth/current [post]
func (h *Handlers) CurrentUser(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	user := userInfo.(*AuthModel.Claims)

	if user == nil {
		log.Error().Err(errors.New("empty user")).Msg("[CurrentUser][User] ")
		response.Error(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	response.JSON(c, http.StatusOK, &dto.CurrentUserResponse{Nim: user.Nim})
}
