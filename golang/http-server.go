/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 *
 *
 */

package main

import (
	"log"
	"net"
	"regexp"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type Request struct {
	Method string
	Html   string
}

func check(err error) {
	if err != nil {
		log.Fatal("[ERROR]\t%v", err)
	}
}

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer l.Close()
	log.Printf("[INFO]\tlistening...\t:%v:%v\n", CONN_HOST, CONN_PORT)

	for {
		conn, err := l.Accept()
		check(err)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	check(err)

	//parse request
	contents := string(buf[:reqLen])
	header := strings.Split(contents, "\n")
	request := parseRequest(header[0])

	log.Printf("[INFO]\t\tmethod       :%v\n", request.Method)
	log.Printf("[INFO]\t\tpage         :%v\n", request.Html)

	//TODO parse json object
	for i, _ := range header[1:] {
		log.Printf("[INFO]\t\treply message:%v\n", header[i+1])
	}

	//TODO: generate response body

	//TODO: reply response
	//conn.Write([]byte("Message received."))
}

func parseRequest(str string) Request {
	reg, _ := regexp.Compile("(?m)([A-Z]+)")

	method := reg.FindString(str)
	html := strings.Replace(str, method+" ", "", 1)
	html = strings.Replace(html, " HTTP/1.1", "", 1)

	return Request{Method: method, Html: html}
}
