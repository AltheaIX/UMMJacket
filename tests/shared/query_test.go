package shared

import (
	"encoding/json"
	"github.com/AltheaIX/UMMJacket/shared/query"
	"testing"
)

type testStruct struct {
	Name string `json:"name" db:"name_test"`
	Zxc  string `json:"zxc" db:"zxc"`
}

func TestBuildUpdateQuery(t *testing.T) {
	var st testStruct

	payload := `{
		"name": "Testttt",
		"zxc": "Testttt",
		"asda": "Testttt"
	}`

	err := json.Unmarshal([]byte(payload), &st)
	if err != nil {
		t.Fatal(err)
	}

	q, params, err := query.BuildUpdateQuery(st, "jackets", 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(q)
	t.Log(params)
}
