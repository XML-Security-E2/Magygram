package service

import (
	"bytes"
	"fmt"
	"html/template"
	"magyAgent/conf"
	"net/smtp"
)

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

var auth smtp.Auth

func SendActivationMail(receiver string, name string, activationId string) {
	auth = smtp.PlainAuth("MagyGram", conf.Current.Mail.Sender, conf.Current.Mail.Password, conf.Current.Mail.Host)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL: "https://" + conf.Current.Server.Host + ":" + conf.Current.Server.Port + "/users/activate/" + activationId,
	}
	r := NewRequest([]string{receiver}, "Hello "+ name + "!", "Hello "+ name + "!")
	if err := r.parseTemplate("mailActivation.html", templateData); err == nil {
		ok, _ := r.sendEmail()
		fmt.Println(ok)
	}
}

func SendResetPasswordMail(receiver string, name string, activationId string) {
	auth = smtp.PlainAuth("MagyGram", conf.Current.Mail.Sender, conf.Current.Mail.Password, conf.Current.Mail.Host)
	templateData := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL: "https://" + conf.Current.Server.Host + ":" + conf.Current.Server.Port + "/users/reset-password/" + activationId,
	}
	r := NewRequest([]string{receiver}, "Hello "+ name + "!", "Hello "+ name + "!")
	if err := r.parseTemplate("mainResetPassword.html", templateData); err == nil {
		ok, _ := r.sendEmail()
		fmt.Println(ok)
	}
}

func (r *Request) sendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := conf.Current.Mail.Host + ":" + conf.Current.Mail.Port

	if err := smtp.SendMail(addr, auth, conf.Current.Mail.Sender, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) parseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
