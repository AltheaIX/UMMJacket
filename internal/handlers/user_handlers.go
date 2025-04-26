package handlers

import (
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handlers) GetUsers(c *gin.Context) {
	var filters filter.Filters
	if err := c.ShouldBindJSON(&filters); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.userService.GetUsersService(c, &filters)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, users)
}
