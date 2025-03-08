package api_test

import (
	"go_crud/api"
	"testing"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Errorf("Error was not returned")
	}
}
