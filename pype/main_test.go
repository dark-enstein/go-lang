package main

import (
	"reflect"
	"testing"
)

type response struct {
	code []int
	err  []error
}

func TestPype(t *testing.T) {
	var codex []int
	pyP := []string{"request", "requests"}
	var expected response
	expected.code = []int{404, 200}

	for _, k := range pyP {

		code, _ := request(k)
		codex = append(codex, code)

	}

	if !reflect.DeepEqual(codex, expected.code) {
		t.Errorf("Test does not pass")
	}

}
