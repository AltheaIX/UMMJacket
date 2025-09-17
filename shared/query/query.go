package query

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"reflect"
	"strings"
)

func BuildUpdateQuery[T any](data T, tableName string, id int) (string, []interface{}, error) {
	query := fmt.Sprintf("UPDATE %v SET ", tableName)
	var clauses []string
	var params []interface{}

	typeOf := reflect.TypeOf(data)
	valueOf := reflect.ValueOf(data)

	if typeOf.Kind() != reflect.Struct {
		err := errors.New("invalid struct")
		log.Err(err).Msg("[Query][BuildUpdateQuery]")
		return "", nil, err
	}

	for i := 0; i < valueOf.NumField(); i++ {
		field := typeOf.Field(i).Tag.Get("db")
		value := valueOf.Field(i)

		if value.Kind() == reflect.Ptr && value.IsNil() {
			continue
		}

		clauses = append(clauses, fmt.Sprintf("%v=?", field))

		if value.Kind() == reflect.Ptr {
			params = append(params, value.Elem().Interface())
		} else {
			params = append(params, value.Interface())
		}
	}

	query += strings.Join(clauses, ", ")
	query += " WHERE id=?"
	params = append(params, id)

	return query, params, nil
}
