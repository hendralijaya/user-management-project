package helper

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendMail(body string, subject string,to string, ccEmail string, ccName string) {
	err := godotenv.Load()
	PanicIfError(err)
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortString := os.Getenv("SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortString)
	PanicIfError(err)
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpSender := os.Getenv("SMTP_SENDER_NAME")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", smtpSender)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", ccEmail, ccName)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)
	dialer := gomail.Dialer{
		Host:     smtpHost,
		Port:     smtpPort,
		Username: smtpUsername,
		Password: smtpPassword,
	}
	err= dialer.DialAndSend(mailer)
    PanicIfError(err)

	log.Println("Mail sent!")
}
