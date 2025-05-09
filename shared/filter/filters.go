package filter

import (
	"errors"
	"fmt"
)

type Filters struct {
	Filter []Filter `json:"filters"`
}

type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

const (
	OperatorEQ = "eq"
)

func BuildFilterAnd(filters []Filter, table string) (string, []interface{}, error) {
	var query string

	var params []interface{}

	query += "SELECT * FROM " + table

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
