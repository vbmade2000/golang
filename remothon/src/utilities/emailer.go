// Email sender library

package utilities

import (
	"net/smtp"
	"strconv"
)

// Struct to store mail related data
type Emailer struct {
	To         []string
	From       string
	Subject    string
	Body       string
	SmtpServer string
	Port       int
	Username   string
	Password   string
	Background bool // For future use

}

// Send() sends email according to specified details.
func (emailer Emailer) Send() error {

	auth := smtp.PlainAuth("", emailer.From, emailer.Password, emailer.SmtpServer)
	body := []byte("Subject: " + emailer.Subject + "\r\n\r\n" + emailer.Body)
	err := smtp.SendMail(emailer.SmtpServer+":"+strconv.Itoa(emailer.Port), auth, emailer.From, emailer.To, body)
	return err
}
