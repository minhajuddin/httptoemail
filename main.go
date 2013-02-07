package main

import (
	"log"
	"net/http"
	"net/smtp"
)

func main() {
	log.Println("started")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a request")
		w.Write([]byte("This is cool"))
	})
	http.ListenAndServe(":8080", nil)
}

func sendEmail() {
	smtp.SendMail("localhost:25",
		nil,
		"HTTP to Email <bot@cosmicvent.com>",
		[]string{"minhajuddin@mailinator.com"},
		[]byte("This is awesome"))
}
