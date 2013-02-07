package main

import (
	"database/sql"
	"fmt"
	"github.com/bmizerany/pq"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
)

var (
	STD_RESPONSE = []byte("DONE")
	DB           *sql.DB
)

func main() {
	connectToDB()
	http.HandleFunc("/", requestHandler)
	err := http.ListenAndServe(":8080", nil)
	handleError(err, "HTTP SERVER")
}

func connectToDB() {
	var err error
	dataSource, err := pq.ParseURL(os.Getenv("PGURL"))
	handleError(err, "URL PARSE")
	DB, err = sql.Open("postgres", dataSource)
	handleError(err, "DB CONNECT")
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

	out(u.Path)

	var s sql.NullString
	hash := "MINK"
	err := DB.QueryRow("SELECT email FROM emails WHERE hash = $1", hash).Scan(&s)
	if err != nil {
		log.Println("SCAN ERR", err)
		return ""
	}
	if !s.Valid {
		out("Email for not found for : " + hash)
		return ""
	}
	return s.String
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

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
func out(args ...interface{}) {
	log.Println(args...)
}
