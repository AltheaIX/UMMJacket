package dto

import (
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"math"
)

type Metadata struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
}

func MetadataFromFilters(filter filter.Filters, totalData int) Metadata {
	return Metadata{
		Page:      filter.Pagination.Page,
		PageSize:  filter.Pagination.PageSize,
		TotalPage: int(math.Ceil(float64(totalData) / float64(filter.Pagination.PageSize))),
		TotalData: totalData,
	}
}
