/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 *
 * refs: http://www.freefavicon.com/freefavicons/objects/iconinfo/yin-yang-152-1271.html
 *
 * refs: http://tools.ietf.org/html/rfc2616#section-9.6
 */

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
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
		log.Fatal(err.Error())
		return
	}

	// 解析処理
	contents := string(buf[:reqLen])
	request := parseRequest(contents, cfg.Server.DocumentRoot)
	println(contents)

	msg := fmt.Sprintf("[INFO]\t\tmethod: %v\t action: %v",
		request.Method, request.Html)
	printOut(msg, green, nil)

	switch request.Method {
	case "GET":
		responseGetMethod(conn, request)

	case "POST":
		//TODO: generate response body
		msg := "[INFO]\t\t:go to post action."
		printOut(msg, green, nil)
		responsePostMethod(conn, request)

	case "PUT":
		//TODO: generate response body
		responsePutMethod(conn, request)

	case "DELETE":
		responseDeleteMethod(conn, request)

	default:
		responseGetMethod(conn, request)
	}
	conn.Close()
	io.Copy(conn, conn)
}

/**
 *  GET要求への処理
 *
 * TODO:
 *   page要求とfavicon要求が来た際に、
 *   両方を返却してブラウザ編集が継続できるように改修
 *
 */
func responseGetMethod(conn net.Conn, request Request) {
	htmlData, err := ioutil.ReadFile(request.Path)
	check(err)

	conn.Write(htmlData)
}

/**
 * * POST要求への処理
 *
 * TODO: 存在しないaction指定の場合は、RoutingErrorを返却させる
 * TODO: multiForm対応させる
 *
 */
func responsePostMethod(conn net.Conn, request Request) {
	htmlData, err := ioutil.ReadFile(request.Path)
	check(err)

	conn.Write(htmlData)
}

/**
 * PUT要求への処理
 *
 * リソースが存在しない場合は新規で作成し、
 * リソースが存在する場合は上書き実施
 *
 */
func responsePutMethod(conn net.Conn, request Request) {
	//TODO Bodyを受け取る口を用意（ファイルアップロードとする）
	ioutil.WriteFile(request.Path, []byte(request.Body), 0644)

	returnStatus := "204"
	conn.Write([]byte(returnStatus))
}

/**
 * DELETE要求への処理
 *
 * リソースが存在する場合は削除する
 *
 */
func responseDeleteMethod(conn net.Conn, request Request) {
	_, err := os.Lstat(request.Path)
	check(err)

	var returnStatus string
	if os.IsNotExist(err) {
		msg := request.Path + " not found"
		printOut(msg, yellow, nil)

		returnStatus = "202"
	} else {
		err := os.Remove(request.Path)
		check(err)

		returnStatus = "204"
	}

	conn.Write([]byte(returnStatus))
}
