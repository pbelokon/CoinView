package api_test

import (
	"coinView/go/crypto/api"
	"testing"
)

// The _test suffix in the package name allows you to write tests as an external user of the package. To test unexported functions you can just use the package name.

func TesAPICall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil { 
		t.Error("Error was not found")
	}
}