package main

import (
	"./utilities"
	"fmt"
)

func main() {
	//IPs, err := utilities.GetIPs()
	emailer := utilities.Emailer{}
	emailer.To = []string{"vbmade2000@gmail.com"}
	emailer.From = "mlvora.2010@gmail.com"
	emailer.Subject = "test subject"
	emailer.Body = "This is stest body"
	emailer.SmtpServer = "smtp.gmail.com"
	emailer.Port = 587
	emailer.Username = "mlvora.2010@gmail.com"
	emailer.Password = "meteor001"
	fmt.Println(emailer.To)
	emailErr := emailer.Send()
	if emailErr != nil {
		fmt.Println("Error : ", emailErr)
	} else {
		fmt.Println("Mail is sent successfully")
	}
	//utilities.StartFileServer()
}
