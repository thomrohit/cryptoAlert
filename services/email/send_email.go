package email

import (
	"net/smtp"
	"os"

	"github.com/labstack/gommon/log"
)

func SendEmail(msg string) {
	// from is senders email address

	// we used environment variables to load the
	// email address and the password from the shell
	// you can also directly assign the email address
	// and the password
	from := os.Getenv("MAIL_USER")
	password := os.Getenv("MAIL_PASSWORD")

	// toList is list of email address that email is to be sent.
	toList := []string{os.Getenv("EMAIL")}

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := "smtp.mailtrap.io"

	// Its the default port of smtp server
	port := "2525"

	// We can't send strings directly in mail,
	// strings need to be converted into slice bytes
	body := []byte("Subject: Price alert\r\n\r\n" + msg)

	// PlainAuth uses the given username and password to
	// authenticate to host and act as identity.
	// Usually identity should be the empty string,
	// to act as username.
	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	// The email is sent to all address in the toList,
	// the body should be of type bytes, not strings
	// This returns error if any occurred.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		log.Fatal("Sending Email failed:", err)
		os.Exit(1)
	}
	log.Info("Successfully sent mail to all user in toList")

}
