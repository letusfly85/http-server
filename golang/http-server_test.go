package main

import (
	"log"
	"net"
	"testing"
	"time"
)

//TODO
func TestResponseGetMethod(t *testing.T) {
	log.Printf("start")

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)

	ch := make(chan string)
	var conn net.Conn
	go func(c chan string) {
		for {
			conn, err := l.Accept()
			check(err)

			request := Request{}
			request.Method = "GET"
			request.Path = "test/get_test.html"
			responseGetMethod(conn, request)
			buf := make([]byte, 1024)
			rlen, err := conn.Read(buf)
			log.Printf(string(buf[:rlen]))
			c <- string(buf[:rlen])
		}
		defer conn.Close()
		defer l.Close()
	}(ch)

	//TODO goroutineの完了条件を付与して、以下のスリープ処理は削除する
	time.Sleep(2 * time.Second)

	//TODO 受け取ったchの値をstringに変換して、assertを実行する
	log.Print(ch)
	log.Printf("finish")

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
