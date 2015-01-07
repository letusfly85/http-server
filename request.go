package main

type Request struct {
	Method string
	Html   string
	Path   string
	Params map[string]string
	Body   string
}
