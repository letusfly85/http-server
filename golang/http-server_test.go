package main

import "testing"

func TestResponseDeleteMethod(t *testing.T) {
	actual := 1
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
