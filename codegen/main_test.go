package main

import (
	"os"
	"testing"
)

func whenFlagHasNoString(t *testing.T) {
	expectedTypeNostring := false
	os.Args = []string{"codegen", "--type", "qr", "--string", "none", "dest", "/users/gent/mane"}
	request := flagParse()
	if request.Present != false {
		actualTypeNostring := request.Present
		t.Errorf("Expected %v got %v", expectedTypeNostring, actualTypeNostring)
	}
}

func TestFlagsParsing(t *testing.T) {
	whenFlagHasNoString(t)
}
