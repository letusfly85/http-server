package main

import (
	"io"
	"net"
	"os"
	"strings"
	"testing"
)

/**
 *	NOTE:
 *   現状、テスト実行前にはhttp_server.goを実行する必要があるが
 *   goroutineなどで、非同期で裏で動かしてテストを実行できるようにしたい
 *
 */
func TestResponseGetMethod(t *testing.T) {
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer conn.Close()

	ch := make(chan string)
	go func(c chan string) {
		// var/www/myhtml/ の直下にファイルを格納している
		// シンボリック対応している
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

/**
 * 削除対象ファイルをテストメソッド中で準備し、
 * deleteメソッド実施後に、存在有無チェックを行う
 *
 */
func TestResponseDeleteMethod(t *testing.T) {
	//delete 対象のファイル準備を実施
	//NOTE: setupメソッドのようなものを提供するライブラリがあればそちらに切り替え
	srcFileName := "test/delete_test_template.html"
	dstFileName := "/var/www/myhtml/test/delete_test.html"
	src, _ := os.Open(srcFileName)
	dst, _ := os.Create(dstFileName)
	_, err := io.Copy(src, dst)
	check(err)

	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer conn.Close()

	ch := make(chan string)
	go func(c chan string) {
		// var/www/myhtml/ の直下にファイルを格納している。
		// シンボリック対応している。
		str := "DELETE /test/delete_test.html HTTP/1.1\n\n\n"
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

	expected := "204"
	if actual != expected {
		t.Errorf("got %vwant %v", actual, expected)
	}

	_, _ = os.Stat(dstFileName)
	if os.IsExist(err) {
		t.Errorf("file not deleted")
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
