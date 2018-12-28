package main

import "log"

func LogData(severity string, message string) {

	log.SetPrefix("Junius ")
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Printf(severity + ": " + message)
}

func main() {
	LogData("error", "panic")
}
