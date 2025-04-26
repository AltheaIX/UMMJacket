package shared

import (
	"encoding/json"
	"github.com/AltheaIX/UMMJacket/shared/filter"
	"testing"
)

func TestJsonFilterPayload(t *testing.T) {
	var filters filter.Filters

	rawJson := []byte(`{
	  "filters": [
		{
		  "field": "nim",
		  "operator": "eq",
		  "value": "test"
		}
	  ]
	}`)

	err := json.Unmarshal(rawJson, &filters)
	if err != nil {
		t.Error(err)
	}

	t.Log(filters.Filter)

	query, params, err := filter.BuildFilterAnd(filters.Filter, "users")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(query)
	t.Log(params)
}
