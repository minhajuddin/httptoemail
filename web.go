package main

import (
	"log"
	"net/http"
)

var (
	STD_RESPONSE = []byte("DONE")
)

func initRoutes() {
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":8080", nil)
	handleError(err, "HTTP SERVER")
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("processing request")
	subject := r.FormValue("subject")
	body := r.FormValue("body")
	to := getReceiver(r.URL)
	go sendEmail(to, subject, body)
	w.Write(STD_RESPONSE)
}
