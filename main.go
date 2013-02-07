package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
)

var STD_RESPONSE = []byte("DONE")

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("processing request")
	subject := r.FormValue("subject")
	body := r.FormValue("body")
	to := getReceiver(r.URL)
	go sendEmail(to, subject, body)
	w.Write(STD_RESPONSE)
}

func getReceiver(u *url.URL) string {
	//TODO: make sure that it is verified
	log.Println(u.Path)
	return "minhajuddin@mailinator.com"
}

func sendEmail(to, subject, body string) {
	msg := fmt.Sprintf(`From: HTTP to Email <bot@cosmicvent.com>
To: %s
Subject: %s

%s`, to, subject, body)
	smtp.SendMail("localhost:25",
		nil,
		"HTTP to Email <bot@cosmicvent.com>",
		[]string{to},
		[]byte(msg))
}
