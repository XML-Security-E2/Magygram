package service

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var BASE_URL = "https://localhost:443/users/activate/"

var auth smtp.Auth


//SREDITI OVO SVEEE
func SendMessage(email string, name string, activationId string) {
	auth = smtp.PlainAuth("MagyGram", "info.magygram@gmail.com", "magygramxml", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
	}{
		Name: name,
		URL:  BASE_URL + activationId,
	}
	r := NewRequest([]string{email}, "Hello "+ name + "!", "Hello, World!")
	if err := r.ParseTemplate("mailActivation.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}

}

//Request struct
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

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "info.magygram@gmail.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
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
