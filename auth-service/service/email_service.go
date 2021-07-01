package service

import (
	"auth-service/conf"
	"bytes"
	"html/template"
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

func SendAgentTokenMail(receiver string, token string) {
	auth = smtp.PlainAuth("MagyGram", conf.Current.Mail.Sender, conf.Current.Mail.Password, conf.Current.Mail.Host)
	templateData := struct {
		Token  string
	}{
		Token: token,
	}
	r := NewRequest([]string{receiver}, "Hello agent" + "!", "Hello agent" + "!")
	if err := r.parseTemplate("agentTokenMain.html", templateData); err == nil {
		if ok, _ := r.sendEmail(); ok {
		} else {
		}
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