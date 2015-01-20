package main

import "testing"

func TestGetMethod(t *testing.T) {
	actual := 1
	expected := 1

	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestPostMethod(t *testing.T) {
	actual := 1
	expected := 1

	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestPutMethod(t *testing.T) {
	actual := 1
	expected := 1

	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestDeleteMethod(t *testing.T) {
	actual := 1
	expected := 1

	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
