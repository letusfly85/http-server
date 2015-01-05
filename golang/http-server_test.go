package main

import (
	"net"
	"strings"
	"testing"
)

/*
	NOTE:
		現状、テスト実行前にはhttp_server.goを実行する必要があるが
		goroutineなどで、非同期で裏で動かしてテストを実行できるようにしたい
*/
func TestResponseGetMethod(t *testing.T) {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer conn.Close()

	ch := make(chan string)
	go func(c chan string) {
		//TODO /var/www/myhtml/ の直下にファイルを格納している。
		//     シンボリック対応ができるように改修する
		str := "GET /test/get_test2.html HTTP/1.1\n\n\n"
		conn.Write([]byte(str))

		buf := make([]byte, 1024)
		rlen, err := conn.Read(buf)
		if err != nil {
			c <- "something happen!"
		}
		c <- string(buf[:rlen])
	}(ch)
	actual := <-ch
	//TODO 改行が含まれてしまっているので対処が必要
	actual = strings.Trim(actual, "\n")

	expected := "hello, test"
	if actual != expected {
		t.Errorf("got %vwant %v", actual, expected)
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
