package handlers

import (
	"github.com/AltheaIX/UMMJacket/internal/domain/auth/model"
	"github.com/AltheaIX/UMMJacket/internal/domain/transaction/model/dto"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

// NewTransaction
// @Summary      Create new transaction
// @Description  Endpoint to make a new transaction
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Param        request body dto.NewTransactionRequest true "Request"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/transactions/new [post]
func (h *Handlers) NewTransaction(c *gin.Context) {
	userInfoCtx, _ := c.Get("userInfo")
	userInfo := userInfoCtx.(*model.Claims)
	userId, err := strconv.Atoi(userInfo.Subject)

	if err != nil {
		log.Error().Err(err).Msg("[NewTransaction][Atoi]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	var request dto.NewTransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	lastId, err := h.transactionService.NewTransactionServices(c, userId, request)
	if err != nil {
		log.Error().Err(err).Msg("[NewTransaction][NewTransactionServices]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	var res dto.NewTransactionResponse
	copier.Copy(&res, &request)
	res.UserId = userId
	res.Id = lastId

	response.JSON(c, 200, res)
}
