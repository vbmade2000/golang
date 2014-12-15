package utilities

import (
        "strconv"
        "net/smtp"
)

/*
   Use this struct to provide required data for email like To address, From address
   Subject and Message 
*/
type EmailData struct {
   To []string
   From string
   Subject string   
   Message string
}

/*
   Use this struct to provide required data for send email like SMTP Server address, SMTP Server port, 
   Username and Password
*/
type ServerData struct {
   Server string
   Port int
   Username string
   Password string
}

/*
   Function to send email
*/
func SendEmail(emailData EmailData, serverData ServerData)(error) {
   pwd := serverData.Password
   body := []byte("Subject: " + emailData.Subject + "\r\n\r\n" + emailData.Message)
   auth := smtp.PlainAuth("", emailData.From, pwd, serverData.Server)
   port := strconv.Itoa(serverData.Port)
   err := smtp.SendMail(serverData.Server + ":" + port, auth, emailData.From, emailData.To, body)
   return err
}

