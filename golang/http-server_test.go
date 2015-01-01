package main

import (
	"net"
	"testing"
)

//TODO
func TestResponseGetMethod(t *testing.T) {
	//TODO: exchange to dummy sever
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer l.Close()

	/*
		conn, err := l.Accept()
		defer conn.Close()
		check(err)

		request := Request{}
		request.Method = "GET"
		request.Path = "test/get_test.html"
		print("start!")
		go responseGetMethod(conn, request)
		print("finish!")
	*/

	actual := 1
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//TODO
func TestResponseDeleteMethod(t *testing.T) {
	actual := 1
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//TODO
func TestResponsePueMethod(t *testing.T) {
	actual := 1
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

//TODO
func TestResponsePostMethod(t *testing.T) {
	actual := 1
	expected := 1
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
