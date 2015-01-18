/**
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var cfg Config = Config{}

func main() {
	var confName = flag.String("f", "conf.gcfg", "config file")
	flag.Parse()

	cfg, err := GetConfig(*confName)
	check(err)

	l, err := net.Listen(cfg.Server.ConnectionType,
		cfg.Server.HostName+":"+cfg.Server.PortNumber)
	check(err)
	defer l.Close()

	msg := fmt.Sprintf("[INFO]\t\tlistening...\t%v", cfg.ToString())
	printOut(msg, green, nil)

	for {
		conn, err := l.Accept()
		check(err)

		go handleRequest(conn, cfg)
	}
}

/**
 * クライアントからリクエスト要求を受け取り、レスポンスを返却する
 *  リクエスト処理の読込
 *  リクエスト処理の解析
 *  レスポンス処理の実行
 *
 *
 *
 */
func handleRequest(conn net.Conn, cfg Config) {
	buf := make([]byte, 4096)
	reqLen, err := conn.Read(buf)
	if err != nil {
		if err.Error() == "EOF" {
			log.Println("return")
			conn.Close()
			return
		}
		log.Println("error handleRequest")
		log.Println(err.Error())
	}

	// 解析処理
	contents := string(buf[:reqLen])
	request := parseRequest(contents, cfg.Server.DocumentRoot)
	println(contents)

	msg := fmt.Sprintf("[INFO]\t\tmethod: %v\t action: %v",
		request.Method, request.Html)
	printOut(msg, green, nil)

	var response Response
	switch request.Method {
	case "GET":
		response, err = getMethod(request)

	case "POST":
		//TODO: generate response body
		response, err = postMethod(request)

	case "PUT":
		response, err = putMethod(request)

	case "DELETE":
		response, err = deleteMethod(request)

	default:
		response, err = getMethod(request)
	}
	check(err)
	conn.Write(response.generateResponce())

	defer conn.Close()
}
