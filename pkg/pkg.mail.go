package pkg

import (
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

func SendMail(email string, token string) {
	CONFIG_AUTH_EMAIL := os.Getenv("CONFIG_AUTH_EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("CONFIG_AUTH_PASSWORD")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "Online Coffee Shop <zwallet6@gmail.com>")
	mailer.SetHeader("To", email)
	mailer.SetAddressHeader("Cc", "zwallet6@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Activation Account")
	mailer.SetBody("text/html", "Open this link to activate your account : localhost:8080/auth/"+token)
	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Mail sent")
	return
}
