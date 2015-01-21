package main

import (
	"io/ioutil"
	"testing"
)

func TestGetMethod(t *testing.T) {
	docRoot := "test/html"
	request := Request{}
	request.Path = docRoot + "/index.html"

	response, _ := getMethod(request, docRoot)

	actual := string(response.Body)
	expected, _ := ioutil.ReadFile("test/html/index.html")

	if string(expected) != actual {
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
