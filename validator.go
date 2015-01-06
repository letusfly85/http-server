package main

import (
	"fmt"

	"github.com/fatih/color"
)

var red = color.New(color.FgRed).SprintFunc()

func check(err error) {
	if err != nil {
		msg := fmt.Sprintf("[ERROR]\t\t%v", err.Error())
		printOut(msg, red, err)
	}
}
