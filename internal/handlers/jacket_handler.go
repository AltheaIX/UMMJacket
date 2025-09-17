package handlers

import (
	"github.com/AltheaIX/UMMJacket/internal/domain/jackets/model/dto"
	"github.com/AltheaIX/UMMJacket/shared"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"github.com/AltheaIX/UMMJacket/shared/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

// CreateJacket
// @Summary      Create Jacket Endpoint
// @Description  Endpoint to create a new jacket
// @Tags         Jackets
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Param        request body dto.CreateJacketsRequest true "Request"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/jackets [post]
func (h *Handlers) CreateJacket(c *gin.Context) {
	var request dto.CreateJacketsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[FiltersJacket][ShouldBindJSON]")
		err = &shared.AppError{Code: http.StatusBadRequest, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	lastId, err := h.jacketsService.InsertJacketsServices(c, request)
	if err != nil {
		log.Error().Err(err).Msg("[CreateJacket][InsertJacketsServices]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	var res dto.CreateJacketsResponse
	copier.Copy(&res, &request)
	res.Id = lastId

	response.JSON(c, 200, res)
}

// FiltersJacket
// @Summary      Filter Jacket Endpoint
// @Description  Endpoint to filter jacket
// @Tags         Jackets
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Param        request body filter.Filters true "Filter"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/jackets/filter [post]
func (h *Handlers) FiltersJacket(c *gin.Context) {
	var request filter.Filters
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[FiltersJacket][ShouldBindJSON]")
		err = &shared.AppError{Code: http.StatusBadRequest, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	jackets, err := h.jacketsService.ResolveJacketsServices(c, &request)
	if err != nil {
		log.Error().Err(err).Msg("[FiltersJacket][ResolveJacketsServices]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	totalData := jackets[0].TotalData
	metadata := dto.MetadataFromFilters(request, totalData)

	response.Metadata(
		c,
		jackets,
		metadata,
	)
}

// UpdateJacket
// @Summary      Update Jacket Endpoint
// @Description  Endpoint to update a jacket
// @Tags         Jackets
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Param id path string true "The Jackets identifier."
// @Param request body dto.UpdateJacketsRequest true "request"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/jackets/{id} [put]
func (h *Handlers) UpdateJacket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil || id == 0 {
		err := &shared.AppError{Code: http.StatusBadRequest, Message: "Invalid ID"}
		response.Error(c, shared.GetCode(err), err.Error())
		return
	}

	var request dto.UpdateJacketsRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error().Err(err).Msg("[UpdateJacket][ShouldBindJSON]")
		err = &shared.AppError{Code: http.StatusBadRequest, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	rowsAffected, err := h.jacketsService.UpdateJacketsServices(c, request, id)
	if err != nil {
		log.Error().Err(err).Msg("[UpdateJacket][UpdateJacketsServices]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	if rowsAffected == 0 {
		err := &shared.AppError{Code: http.StatusNotFound, Message: "Data not found"}
		response.Error(c, shared.GetCode(err), err.Error())
		return
	}

	response.JSON(c, 200, "Success")
}

// DeleteJacket
// @Summary      Delete Jacket Endpoint
// @Description  Endpoint to delete a jacket
// @Tags         Jackets
// @Accept       json
// @Produce      json
// @Param Authorization header string false "<User Authorization>"
// @Param id path string true "The Jackets identifier."
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      401  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /v1/jackets/{id} [delete]
func (h *Handlers) DeleteJacket(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		err := &shared.AppError{Code: http.StatusBadRequest, Message: "Invalid ID"}
		response.Error(c, shared.GetCode(err), err.Error())
		return
	}

	rowsAffected, err := h.jacketsService.DeleteJacketsServices(c, id)
	if err != nil {
		log.Error().Err(err).Msg("[DeleteJacket][DeleteJacketsServices]")
		err = &shared.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		code := shared.GetCode(err)
		response.Error(c, code, err.Error())
		return
	}

	if rowsAffected == 0 {
		err := &shared.AppError{Code: http.StatusNotFound, Message: "Data not found"}
		response.Error(c, shared.GetCode(err), err.Error())
		return
	}

	response.JSON(c, 200, dto.DeleteJacketsResponse{Deleted: true})
}
