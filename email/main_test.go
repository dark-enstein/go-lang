package main

import (
	"reflect"
	"testing"
)

func TestEmailValidate(t *testing.T) {
	emails := []string{"ayobamafu34@gmail.com", "handle45", "3894#*$H#@knockoff.com", "noone.com"}
	expected := []bool{true, false, false, false}

	_, actual := emailValidate(emails)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
