package email

import (
	"bytes"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type AppointmentData struct {
	Name          string
	Date          string
	Time          string
	Motivo        string
	AppointmentID int
}

func SendEmail(to string, subject string, data AppointmentData) error {

	template, err := template.ParseFiles("pages/email.html")

	if err != nil {
		log.Println("Failed to parse email template", err)
		return err
	}

	var body bytes.Buffer
	if err := template.Execute(&body, data); err != nil {
		log.Panicln("Failed to execute email template:")
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "acneclinic2024@outlook.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp-mail.outlook.com", 587, "acneclinic2024@outlook.com", "Repeatrave1")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil

}
