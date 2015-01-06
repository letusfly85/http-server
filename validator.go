package main

import "fmt"

func check(err error) {
	if err != nil {
		msg := fmt.Sprintf("[ERROR]\t\t%v", err.Error())
		printOut(msg, red, err)
	}
}
