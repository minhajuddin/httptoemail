package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func sendEmail(to, subject, body string) {
	if len(to) == 0 || !strings.Contains(to, "@") {
		out("No 'to' found, ignoring request :", to)
		return
	}

	out("SENDING EMAIL to :", to)
	msg := fmt.Sprintf(`From: HTTP to Email <bot@minhajuddin.com>
To: %s
Subject: %s

%s`, to, subject, body)
	smtp.SendMail("localhost:25",
		nil,
		"HTTP to Email <bot@minhajuddin.com>",
		[]string{to},
		[]byte(msg))
}
