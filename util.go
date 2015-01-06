package main

import "log"

func printOut(msg string, f func(a ...interface{}) string, err error) {
	if err != nil {
		log.Panicln(err.Error())

	} else {
		log.Printf(f(msg))
	}
}
