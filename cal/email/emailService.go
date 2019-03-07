package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"	
)
  
type to struct{
	Name,Address string
}

func SendEmail(name,email,subject,message string){
	os.Setenv("serverpassrd","")
	reciever := new(to)
	reciever.Name=name
	reciever.Address=email
	SendMail(*reciever,subject,message,"calendar1.ics")
}

func SendMail(reciever struct {Name string; Address string}, subject string, message string, filePath string) {

	d := gomail.NewDialer("smtp.gmail.com", 25, "", os.Getenv("serverpassrd"))
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	m := gomail.NewMessage()
	m.SetHeader("From", "")
	m.SetAddressHeader("To", reciever.Address, reciever.Name)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", fmt.Sprintf("Hello %s!", reciever.Name) + "\n " + message)
	m.Attach(filePath)
	if err := gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %q: %v", reciever.Address, err)
	}
	m.Reset()

}
