package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseRequest(contents string, documentRoot string) Request {
	var request = Request{}

	target := strings.Split(contents, "\n")
	header := target[0]
	params := target[len(target)-1]

	request.parseHeader(header)
	request.setRequestPath(documentRoot)
	request.parseFormParams(params)

	return request
}

func (request *Request) parseHeader(header string) {
	//TODO 引数チェックの実装

	reg4method, _ := regexp.Compile("(?m)([A-Z]+)")
	method := reg4method.FindString(header)
	html := strings.Replace(header, method+" ", "", 1)

	reg4version := regexp.MustCompile("HTTP/([0-9]).([0-9])")
	html = reg4version.ReplaceAllString(html, "")

	reg4params := regexp.MustCompile("([?])(.*)")
	html = reg4params.ReplaceAllString(html, "")
	html = strings.TrimSpace(html)

	request.Html = html
	request.Method = method
}

/**
 * Request構造体を受け取り、htmlの絶対パスを設定する
 * リソースの指定がない場合は、indexをデフォルトページとして返却する
 * 指定されたリソースが存在しない場合、404エラーを返却する
 * PUT処理の場合はリソースを新規作成するために、404は返却しない
 *
 */
func (request *Request) setRequestPath(documentRoot string) {
	if request.Html == "/" || request.Html == "" {
		request.Html = "/index.html"

		request.Path = documentRoot + request.Html
		return
	}

	path := documentRoot + request.Html

	info, err := os.Lstat(path)

	//ファイルの存在チェック
	if os.IsNotExist(err) && request.Method != "PUT" {
		msg := fmt.Sprintf("[WARN]\t\t%v\n", err)
		printOut(msg, yellow, nil)

		request.Path = documentRoot + "/404.html"
		return
	}

	//ファイルのシンボリックリンクチェック
	if info.Mode()&os.ModeSymlink == os.ModeSymlink {
		linkPath, err := os.Readlink(path)
		check(err)

		request.Path = linkPath
		return

	} else {
		request.Path = path
	}
}

/**
 * リクエスト処理の文字列の最終行をターゲットとしてparameter特定する
 * パラメータをkey, valueの形に分解して、mapに格納する
 * mapはRequestのParamsに格納する
 *
 * TODO:
 *  httpのリクエスト処理の使用上、paramterが最終行以外でも定義可能か確認
 *
 *	TODO urlにパラメータ記載がある場合の処理を実装
 *
 */
func (request *Request) parseFormParams(str string) {
	if str == "" {
		return
	}
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
