package main

import "testing"

//TODO
func TestParseHeader(t *testing.T) {
	expected := 1
	actual := 1

	if expected != actual {
		t.Errorf("got %vwant %v", actual, expected)
	}
}
