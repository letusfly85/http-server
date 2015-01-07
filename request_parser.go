package main

import (
	"regexp"
	"strings"
)

func (request *Request) parseHeader(header string) {
	reg4method, _ := regexp.Compile("(?m)([A-Z]+)")
	method := reg4method.FindString(header)
	html := strings.Replace(header, method+" ", "", 1)

	//TODO regexp replaceに変更する。HTTPのversionは、HTTP/x.xとして表現されるため
	html = strings.Replace(html, " HTTP/1.1", "", 1)
	html = strings.TrimSpace(html)

	request.Html = html
	request.Method = method
}
