/**
 * refs: http://whispering-gophers.appspot.com/talk.slide#16
 * refs: https://coderwall.com/p/wohavg
 * refs: http://d.hatena.ne.jp/taknb2nch/20140210/1392044307
 *
 * refs: http://www.freefavicon.com/freefavicons/objects/iconinfo/yin-yang-152-1271.html
 *
 */

package main

//TODO https://github.com/op/go-logging を利用する
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
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
	Params map[string]string
}

var green = color.New(color.FgGreen, color.Bold).Add(color.Underline).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func printOut(msg string, f func(a ...interface{}) string, err error) {
	if err != nil {
		log.Fatal(f(msg))

	} else {
		log.Printf(f(msg))
	}
}

func check(err error) {
	if err != nil {
		msg := fmt.Sprintf("[ERROR]\t%v", err)
		printOut(msg, red, err)
	}
}

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer l.Close()
	msg := fmt.Sprintf("[INFO]\t\tlistening...\t%v:%v", CONN_HOST, CONN_PORT)
	printOut(msg, green, nil)

	for {
		conn, err := l.Accept()
		check(err)

		go handleRequest(conn)
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
func handleRequest(conn net.Conn) {
	buf := make([]byte, 4096)

	reqLen, err := conn.Read(buf)
	check(err)

	// 解析処理
	contents := string(buf[:reqLen])
	header := strings.Split(contents, "\n")
	request := parseRequest(header[0])
	request = parseForm(header[len(header)-1], request)

	msg := fmt.Sprintf("[INFO]\t\tmethod: %v\t action: %v",
		request.Method, request.Html)
	printOut(msg, green, nil)

	switch request.Method {
	case "GET":
		responseGetMethod(conn, request)

	case "POST":
		//TODO: generate response body
		log.Printf("[INFO]\t\t:%v\n", yellow("go to post action."))
		responsePostMethod(conn, request)

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
	conn.Close()
	io.Copy(conn, conn)
}

/**
 * * クライアントからの要求を、以下のフォーマットに分解する
 * ** Request {Method:string, Html:string}
 * ** json object: Request以外のjson形式で定義されたデータ構造
 *
 *
 */
func parseRequest(str string) Request {
	reg4method, _ := regexp.Compile("(?m)([A-Z]+)")
	method := reg4method.FindString(str)
	html := strings.Replace(str, method+" ", "", 1)
	html = strings.Replace(html, " HTTP/1.1", "", 1)
	html = strings.TrimSpace(html)

	return Request{Method: method, Html: html}
}

/**
 * パラメータをkey, valueの形に分解して、mapに格納する
 * mapはRequestのParamsに格納する
 *
 */
func parseForm(str string, req Request) Request {
	params := make(map[string]string)

	conditions := strings.Split(str, "&")
	for _, condition := range conditions {
		if condition == "" {
			continue
		}
		reg4param := regexp.MustCompile("(.*)=(.*)")
		group := reg4param.FindSubmatch([]byte(condition))
		key, val := string(group[1]), string(group[2])
		params[key] = val
	}
	req.Params = params

	return req
}

/**
 * * GET要求への処理
 * ** ページ指定がない場合は、index.htmlをデフォルトページとして返却する
 *
 * TODO: page要求とfavicon要求が来た際に、両方を返却してブラウザ編集が継続できるように改修
 *
 */
func responseGetMethod(conn net.Conn, request Request) {
	if request.Html == "/" {
		request.Html = "/index.html"
	}

	path := HTML_DIR + request.Html
	if _, err := os.Stat(path); os.IsNotExist(err) {
		msg := fmt.Sprintf("[WARN]\t\t%v\n", err)
		printOut(msg, yellow, nil)
		// ページが存在しない場合は404エラーを返却するようにする
		path = HTML_DIR + "/404.html"
	}

	htmlData, err := ioutil.ReadFile(path)
	check(err)
	conn.Write(htmlData)
}

/**
 * * POST要求への処理
 *
 * TODO: 存在しないaction指定の場合は、RoutingErrorを返却させる。
 * TODO: multiForm対応させる。
 *
 */
func responsePostMethod(conn net.Conn, request Request) {
	//TODO: 書き換え
	path := HTML_DIR + request.Html
	if _, err := os.Stat(path); os.IsNotExist(err) {
		msg := fmt.Sprintf("[WARN]\t\t%v\n", err)
		printOut(msg, yellow, nil)
		path = HTML_DIR + "/404.html"
	}

	htmlData, err := ioutil.ReadFile(path)
	check(err)
	conn.Write(htmlData)
}
