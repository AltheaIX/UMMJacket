package filter

import (
	"errors"
	"fmt"
)

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

const (
	OperatorEQ = "eq"
)

type Sort struct {
	Field string `json:"field"`
	Order string `json:"order"`
}

type Filters struct {
	Filter     []Filter   `json:"filters"`
	Pagination Pagination `json:"pagination"`
	Sort       Sort       `json:"sort"`
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

func (p *Pagination) SettleValue() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 5
	}
}

func BuildFilter(filters *Filters, tableName string) (string, []interface{}, error) {
	var query string
	var params []interface{}
	query += fmt.Sprintf("SELECT (SELECT COUNT(*) FROM %s) AS total_data, t.* FROM %s t", tableName, tableName)

	if filters.Filter != nil {
		queryAnd, paramsAnd, err := BuildFilterAnd(filters.Filter)
		if err != nil {
			return "", nil, err
		}

		params = append(params, paramsAnd...)
		query += queryAnd
	}

	if filters.Sort.Field != "" && filters.Sort.Order != "" {
		query += " ORDER BY " + filters.Sort.Field + " " + filters.Sort.Order
	}

	filters.Pagination.SettleValue()

	query += fmt.Sprintf(
		" LIMIT %d OFFSET %d",
		filters.Pagination.PageSize,
		(filters.Pagination.Page-1)*filters.Pagination.PageSize,
	)

	return query, params, nil
}

func BuildFilterAnd(filters []Filter) (string, []interface{}, error) {
	var query string

	var params []interface{}

	for filterIndex, filter := range filters {
		if filter.Field != "" && filterIndex == 0 {
			query += " WHERE "
		}

		switch filter.Operator {
		case OperatorEQ:
			query += fmt.Sprintf("%s = ?", filter.Field)
		default:
			return "", nil, errors.New("invalid filter operator")
		}

		if filterIndex != len(filters)-1 {
			query += " AND "
		}

		params = append(params, filter.Value)
	}

	return query, params, nil
}

func GetMultipleTableCounts(tableName []string) string {
	var query string

	query += "SELECT "

	for idx, table := range tableName {
		query += fmt.Sprintf("(SELECT COUNT(*) FROM %s) as %s_count", table, table)
		if idx != len(tableName)-1 {
			query += ", "
			continue
		}

		query += ";"
	}

	return query
}
