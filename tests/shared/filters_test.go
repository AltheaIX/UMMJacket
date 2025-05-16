package shared

import (
	"encoding/json"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"testing"
)

func TestJsonFilterPayload(t *testing.T) {
	var filters *filter.Filters

	rawJson := []byte(`{
	 "filters": [
		{
		  "field": "nim",
		  "operator": "eq",
		  "value": "test"
		}
	 ],
	 "pagination": {
	    "page": 1
	 },
	 "sort": {
		"field": "id",
		"order": "ASC"
	 }
	}`)

	//rawJson := []byte(`{}`)

	err := json.Unmarshal(rawJson, &filters)
	if err != nil {
		t.Error(err)
	}

	t.Log(filters.Filter)

	query, params, err := filter.BuildFilter(filters, "users")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(query)
	t.Log(params)
}

func TestGetMultipleTableCounts(t *testing.T) {
	query := filter.GetMultipleTableCounts([]string{"users", "transactions"})
	t.Log(query)
}
