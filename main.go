package main

import (
	"log"
)

func main() {
	connectToDB()
	//createReceiver("minhajuddin@mailinator.com")
	//out("DONE")
	initRoutes()
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func out(args ...interface{}) {
	log.Println(args...)
}
