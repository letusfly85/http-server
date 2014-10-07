/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 *
 *
 */

package main

//TODO https://github.com/op/go-logging を利用する
import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"

	HTML_DIR = "/var/www/myhtml"
)

type Request struct {
	Method string
	Html   string
}

func check(err error) {
	if err != nil {
		log.Fatal("[ERROR]\t%v\n", err)
	}
}

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer l.Close()
	log.Printf("[INFO]\tlistening...\t%v:%v\n", CONN_HOST, CONN_PORT)

	for {
		conn, err := l.Accept()
		check(err)

		go handleRequest(conn)
	}
}

/**
 * * クライアントからリクエスト要求を受け取り、レスポンスを返却する
 * ** リクエスト処理の読込
 * ** リクエスト処理の解析
 * ** レスポンス処理の実行
 *
 *
 *
 */
func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	check(err)

	contents := string(buf[:reqLen])
	header := strings.Split(contents, "\n")
	request := parseRequest(header[0])

	log.Printf("[INFO]\t\tmethod       :%v:\n", request.Method)
	log.Printf("[INFO]\t\tpage         :%v:\n", request.Html)

	//TODO parse json object
	for i, _ := range header[1:] {
		log.Printf("[INFO]\t\treply message:%v\n", header[i+1])
	}

	switch request.Method {
	case "GET":
		responseGetMethod(conn, request)

	case "POST":
		//TODO: generate response body
		responseGetMethod(conn, request)

	case "PUT":
		//TODO: generate response body
		responseGetMethod(conn, request)

	case "DELETE":
		//TODO: generate response body
		responseGetMethod(conn, request)

	default:
		//TODO: generate response body
		responseGetMethod(conn, request)
	}
}

/**
 * * クライアントからの要求を、以下のフォーマットに分解する
 * ** Request {Method:string, Html:string}
 * ** json object: Request以外のjson形式で定義されたデータ構造
 *
 *
 */
func parseRequest(str string) Request {
	reg, _ := regexp.Compile("(?m)([A-Z]+)")

	method := reg.FindString(str)
	html := strings.Replace(str, method+" ", "", 1)
	html = strings.Replace(html, " HTTP/1.1", "", 1)
	html = strings.TrimSpace(html)

	return Request{Method: method, Html: html}
}

/**
 * * GET要求への処理
 * ** ページ指定がない場合は、index.htmlをデフォルトページとして返却する
 * todo connを使いまわしているのでio.Copyするように改修する
 *
 *
 *
 */
func responseGetMethod(conn net.Conn, request Request) {
	if request.Html == "/" {
		log.Printf("enter\n")
		request.Html = "/index.html"
	}

	path := HTML_DIR + request.Html
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("[WARN]\t%v\n", err)
		//TODO ページが存在しない場合は404エラーを返却するようにする

	} else {
		htmlData, err := ioutil.ReadFile(path)
		check(err)

		conn.Write(htmlData)
	}
}
