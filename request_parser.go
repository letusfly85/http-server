package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func (request *Request) parseHeader(contents string) {
	header := strings.Split(contents, "\n")[0]

	reg4method, _ := regexp.Compile("(?m)([A-Z]+)")
	method := reg4method.FindString(header)
	html := strings.Replace(header, method+" ", "", 1)

	//TODO regexp replaceに変更する。HTTPのversionは、HTTP/x.xとして表現されるため
	html = strings.Replace(html, " HTTP/1.1", "", 1)
	html = strings.TrimSpace(html)

	request.Html = html
	request.Method = method
}

/**
 *  クライアントからの要求を、解析してRequest構造体を返却する
 *
 *
 */
func (request *Request) parseBody(contents string) {
	request.setRequestPath()

	target := strings.Split(contents, "\n")
	params := target[len(target)-1]
	if params != "" {
		request.setFormParams(params)
	}
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
