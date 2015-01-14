package main

import (
	"io/ioutil"
	"net"
	"os"
)

/**
 *  GET要求への処理
 *
 * TODO:
 *   page要求とfavicon要求が来た際に、
 *   両方を返却してブラウザ編集が継続できるように改修
 *
 */
func getMethod(request Request) (response Response, err error) {
	htmlData, err := ioutil.ReadFile(request.Path)

	response = Response{}
	response.Body = htmlData
	return response, err
	//conn.Write(htmlData)
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
 * TODO: メソッドの引数改修
 *
 */
func responsePutMethod(conn net.Conn, request Request) {
	ioutil.WriteFile(request.Path, []byte(request.Body), 0644)

	returnStatus := "204"
	conn.Write([]byte(returnStatus))
}

/**
 * DELETE要求への処理
 *
 * リソースが存在する場合は削除する
 *
 * TODO: メソッドの引数改修
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
