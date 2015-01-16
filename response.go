package main

import "time"

type Response struct {
	Status string
	Header []byte
	Body   []byte
}

func (response *Response) generateResponce() []byte {
	now := time.Now()

	res := `HTTP1.1 ` + response.Status + `
Date:` + (now.Format(time.RFC3339)) + `
Server:my_golang_server

` + string(response.Body) + `

`

	return []byte(res)
}
