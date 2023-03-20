package modules

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type Smtp struct {
	addr string
	auth smtp.Auth
	mime string
	from string
}

type SmtpOption struct {
	Subject      string
	Target       []string
	TemplateData interface{}
	FileNames    string
}

func NewSmtp() Smtp {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	from := "From: " + os.Getenv("CONFIG_SENDER_NAME") + "\n"
	addr := fmt.Sprintf("%s:%s", os.Getenv("CONFIG_SMTP_HOST"), os.Getenv("CONFIG_SMTP_PORT"))
	auth := smtp.PlainAuth("", os.Getenv("CONFIG_AUTH_EMAIL"), os.Getenv("CONFIG_AUTH_PASSWORD"), os.Getenv("CONFIG_SMTP_HOST"))

	return Smtp{
		addr: addr,
		auth: auth,
		mime: mime,
		from: from,
	}
}

func (s *Smtp) SendMail(option SmtpOption) error {
	t, err := template.ParseFiles(option.FileNames)
	if err != nil {
		return fmt.Errorf("error while parse files: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, option.TemplateData); err != nil {
		return fmt.Errorf("error while execute templates: %s", err.Error())
	}

	emailBody := buf.String()
	subject := "Subject: " + option.Subject + "!\n"
	msg := []byte(s.from + subject + s.mime + "\n" + emailBody)

	if err := smtp.SendMail(s.addr, s.auth, os.Getenv("CONFIG_AUTH_EMAIL"), option.Target, msg); err != nil {
		return err
	}

	return nil
}
