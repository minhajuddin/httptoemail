package main

import (
	"database/sql"
	"github.com/bmizerany/pq"
	"log"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	DB      *sql.DB
	POST_RX = regexp.MustCompile("^[0-9a-f]+$")
)

func connectToDB() {
	var err error
	dataSource, err := pq.ParseURL(os.Getenv("PGURL"))
	handleError(err, "URL PARSE")
	DB, err = sql.Open("postgres", dataSource)
	handleError(err, "DB CONNECT")
}

func getReceiver(u *url.URL) string {
	//TODO: make sure that it is verified

	hash := strings.Trim(strings.ToLower(u.Path), "/ ")

	if !POST_RX.MatchString(hash) {
		return ""
	}

	var s sql.NullString
	err := DB.QueryRow("SELECT email FROM users WHERE id = $1", hash).Scan(&s)
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

func createReceiver(email string) {
	id := hash(email)
	id += strconv.FormatInt(time.Now().Unix(), 10)
	id = hash(email)

	DB.Exec("INSERT INTO users(id, email) VALUES($1, $2)", id, email)
}
