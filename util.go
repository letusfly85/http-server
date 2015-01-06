package main

import (
	"log"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen, color.Bold).Add(color.Underline).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()

func printOut(msg string, f func(a ...interface{}) string, err error) {
	if err != nil {
		log.Panicln(err.Error())

	} else {
		log.Printf(f(msg))
	}
}
