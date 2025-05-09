package handlers

import (
	"github.com/AltheaIX/UMMJacket/internal/domain/statistic/model/dto"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Dashboard godoc
// @Summary      Dashboard Endpoint
// @Description  Endpoint to get dashboard's statistics
// @Tags         Statistics
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/statistics/dashboard [get]
func (h *Handlers) Dashboard(c *gin.Context) {
	statistic, err := h.statisticService.GetDashboardCount(c)
	if err != nil {
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	response.JSON(
		c,
		http.StatusOK,
		&dto.DashboardResponse{UsersCount: statistic.UsersCount, TransactionsCount: statistic.TransactionsCount},
	)
}
