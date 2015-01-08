package main

import (
	"strings"
	"testing"
)

func TestParseHeader(t *testing.T) {
	contents := `GET /index.html HTTP/2.1
		Host: localhost:3333
		Connection: keep-alive
		Accept: */*
		User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 20_10_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36
		Accept-Encoding: gzip, deflate, sdch
		Accept-Language: ja,en-US;q=0.8,en;q=0.6


		`
	target := strings.Split(contents, "\n")
	header := target[0]

	var request = Request{}
	request.parseHeader(header)

	actual := request.Html
	expected := "/index.html"

	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func TestSetRequestPath(t *testing.T) {
	var request = Request{}

	request.Html = ""
	request.setRequestPath("test/html")

	actual := request.Path
	expected := "test/html/index.html"
	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}

	request.Html = "/"
	request.setRequestPath("test/html")

	actual = request.Path
	expected = "test/html/index.html"
	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}

	request.Html = "/dummy.html"
	request.setRequestPath("test/html")

	actual = request.Path
	expected = "test/html/dummy.html"
	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}

	request.Html = "/symdummy.html"
	request.setRequestPath("test/html")

	actual = request.Path
	expected = "test/html/dummy.html"
	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}

	request.Html = "/notExists.html"
	request.setRequestPath("test/html")

	actual = request.Path
	expected = "test/html/404.html"
	if expected != actual {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
