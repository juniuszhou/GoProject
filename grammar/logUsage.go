package main

import "log"

func LogData(severity string, message string) {
	log.Printf(severity + message)
}
