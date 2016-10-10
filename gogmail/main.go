package main

import "net/smtp"
import "log"

func send(body string, to string) {
	from := " @gmail.com"
	password := " "

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"MIME-Version: 1.0" + " \r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: Your messages subject" + "\r\n\r\n" +
		body + "\r\n"

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, password, "smtp.gmail.com"), from, []string{to}, []byte(msg))
	if err != nil {
		log.Printf("Error: %s", err)
		return
	}

	log.Print("message sent")
}

func main() {
	send("<h1>message sent from go.</h1>", " @gmail.com")
}
