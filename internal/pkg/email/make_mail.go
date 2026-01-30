package mailer

import (
	"bytes"
	"os"
	"text/template"
)

func (m Mailer) MakeMail(path string, data any) Mailer {
	readHtml, err := os.ReadFile(path)
	if err != nil {
		m.Error = err
		return m
	}

	tmpl, err := template.New("custom").Parse(string(readHtml))
	if err != nil {
		m.Error = err
		return m
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		m.Error = err
		return m
	}

	m.Body = strMail.String()

	return m
}
