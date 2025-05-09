package shared

import (
	"github.com/AltheaIX/UMMJacket/shared/crypt"
	"testing"
)

func TestGenerateHashPassword(t *testing.T) {
	hash, err := crypt.GenerateHashPassword("Malik")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hash)
}
