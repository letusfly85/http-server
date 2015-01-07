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
	"net"
	"os"
	"regexp"
	"strings"
)

var cfg Config

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
	request := parseRequest(contents)
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
		//TODO: generate response body
		responseGetMethod(conn, request)
	}
	conn.Close()
	io.Copy(conn, conn)
}

/**
 *  クライアントからの要求を、解析してRequest構造体を返却する
 *
 *
 */
func parseRequest(contents string) Request {
	request := Request{}

	target := strings.Split(contents, "\n")

	header := target[0]
	request.parseHeader(header)
	request.setRequestPath()

	params := target[len(target)-1]
	if params != "" {
		request.setFormParams(params)
	}

	return request
}

/**
 * Request構造体を受け取り、htmlの絶対パスを設定する
 * リソースの指定がない場合は、indexをデフォルトページとして返却する
 * 指定されたリソースが存在しない場合、404エラーを返却する
 * PUT処理の場合はリソースを新規作成するために、404は返却しない
 *
 * TODO:
 *  設定ファイルを用意し、DocumentRootを設定出来るように改修する
 *
 */
func (request *Request) setRequestPath() {
	if request.Html == "/" {
		request.Html = "/index.html"
	}

	path := cfg.Server.DocumentRoot + request.Html
	info, err := os.Lstat(path)
	check(err)

	//シンボリックリンクの場合は、Readlink関数を利用して実パスを取得
	//refs: https://groups.google.com/forum/#!topic/golang-nuts/jpsgja5B_Kk
	//refs: http://golang.org/pkg/os/#ModeSymlink
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		path, err = os.Readlink(path)
	}

	if os.IsNotExist(err) && request.Method != "PUT" {
		msg := fmt.Sprintf("[WARN]\t\t%v\n", err)
		printOut(msg, yellow, nil)
		path = cfg.Server.DocumentRoot + "/404.html"
	}

	request.Path = path
}

/**
 * リクエスト処理の文字列の最終行をターゲットとしてparameter特定する
 * パラメータをkey, valueの形に分解して、mapに格納する
 * mapはRequestのParamsに格納する
 *
 * TODO:
 *  httpのリクエスト処理の使用上、paramterが最終行以外でも定義可能か確認
 *
 */
func (request *Request) setFormParams(str string) {
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
	request.Params = params
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
	//TODO 削除対象のファイルの存在チェック
	err := os.Remove(request.Path)
	check(err)

	returnStatus := "204"
	conn.Write([]byte(returnStatus))
}
